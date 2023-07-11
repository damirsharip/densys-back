package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
	"github.com/pkg/errors"
)

// DepartmentCreate creates a new department Service.
func (s *Service) DepartmentCreate(ctx context.Context, arg entity.Department) (entity.Department, error) {
	var department entity.Department

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		arg.ID = uuid.New().String()

		department, err = q.DepartmentCreate(ctx, arg)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.Department{}, errors.Wrap(err, "[Service.DepartmentCreate] failed to create department")
	}

	return department, nil
}

// DepartmentGetAll gets all departments Service.
func (s *Service) DepartmentGetAll(ctx context.Context) ([]entity.Department, error) {
	var departments []entity.Department

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		departments, err = q.DepartmentGetAll(ctx)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Department{}, errors.Wrap(err, "[Service.DepartmentGetAll] failed to get all departments")
	}

	return departments, nil
}
