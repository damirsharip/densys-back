package pgstorage

import (
	"context"

	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/pkg/errors"
)

// ServiceCreate creates a new department.
func (q *Queries) ServiceCreate(ctx context.Context, arg entity.Service) (entity.Service, error) {
	_, err := q.db.Exec(
		ctx,
		serviceCreateSQL,
		arg.ID,
		arg.Name,
		arg.Price,
		arg.ListOfContradictions,
		arg.OtherRelatedInformation,
		arg.Specialization.ID,
	)

	if err != nil {
		return entity.Service{}, errors.Wrap(err, "[queries.ServiceCreate] failed to create department")
	}

	return arg, nil
}

// ServiceGetAll gets all departments.
func (q *Queries) ServiceGetAll(ctx context.Context) ([]entity.Service, error) {
	rows, err := q.db.Query(ctx, serviceGetAllSQL)

	if err != nil {
		return []entity.Service{}, errors.Wrap(err, "[queries.ServiceGetAll] failed to get all departments")
	}

	defer rows.Close()

	var services []entity.Service

	for rows.Next() {
		var s entity.Service

		err = rows.Scan(
			&s.ID,
			&s.Name,
			&s.Price,
			&s.ListOfContradictions,
			&s.OtherRelatedInformation,
			&s.Specialization.ID,
			&s.Specialization.Name,
		)

		if err != nil {
			return []entity.Service{}, errors.Wrap(err, "[queries.ServiceGetAll] failed to scan department")
		}

		services = append(services, s)
	}

	return services, nil
}

// DoctorServiceGetAll gets all departments.
func (q *Queries) DoctorServiceGetAll(ctx context.Context, userId string) ([]entity.Service, error) {
	rows, err := q.db.Query(ctx, doctorServiceGetAllSQL, userId)

	if err != nil {
		return []entity.Service{}, errors.Wrap(err, "[queries.DoctorServiceGetAll] failed to get all departments")
	}

	defer rows.Close()

	var services []entity.Service

	for rows.Next() {
		var s entity.Service

		err = rows.Scan(
			&s.ID,
			&s.Name,
			&s.Price,
			&s.ListOfContradictions,
			&s.OtherRelatedInformation,
		)

		if err != nil {
			return []entity.Service{}, errors.Wrap(err, "[queries.DoctorServiceGetAll] failed to scan department")
		}

		services = append(services, s)
	}

	return services, nil
}

func (q *Queries) ServiceAutoComplete(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, serviceAutoCompleteSQL)

	if err != nil {
		return []string{}, errors.Wrap(err, "[queries.ServiceAutoComplete] failed to get all departments")
	}

	defer rows.Close()

	var services []string

	for rows.Next() {
		var s string

		err = rows.Scan(
			&s,
		)

		if err != nil {
			return []string{}, errors.Wrap(err, "[queries.ServiceAutoComplete] failed to scan department")
		}

		services = append(services, s)
	}

	return services, nil

}