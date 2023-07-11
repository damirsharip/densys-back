package pgstorage

import (
	"context"

	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/pkg/errors"
)

// DepartmentCreate creates a new department.
func (q *Queries) DepartmentCreate(ctx context.Context, arg entity.Department) (entity.Department, error) {
	_, err := q.db.Exec(
		ctx,
		departmentCreateSQL,
		arg.ID,
		arg.Name,
	)

	if err != nil {
		return entity.Department{}, errors.Wrap(err, "[queries.DepartmentCreate] failed to create department")
	}

	return arg, nil
}

// DepartmentGetAll gets all departments.
func (q *Queries) DepartmentGetAll(ctx context.Context) ([]entity.Department, error) {
	rows, err := q.db.Query(ctx, departmentGetAllSQL)

	if err != nil {
		return []entity.Department{}, errors.Wrap(err, "[queries.DepartmentGetAll] failed to get all departments")
	}

	defer rows.Close()

	var departments []entity.Department

	for rows.Next() {
		var department entity.Department

		err = rows.Scan(
			&department.ID,
			&department.Name,
		)

		if err != nil {
			return []entity.Department{}, errors.Wrap(err, "[queries.DepartmentGetAll] failed to scan department")
		}

		departments = append(departments, department)
	}

	return departments, nil
}
