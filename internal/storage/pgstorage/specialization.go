package pgstorage

import (
	"context"

	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/pkg/errors"
)

// SpecializationCreate creates a new specialization.
func (q *Queries) SpecializationCreate(ctx context.Context, arg entity.Specialization) (entity.Specialization, error) {
	_, err := q.db.Exec(
		ctx,
		specializationCreateSQL,
		arg.ID,
		arg.Name,
	)

	if err != nil {
		return entity.Specialization{}, errors.Wrap(err, "[queries.SpecializationCreate] failed to create specialization")
	}

	return arg, nil
}

// SpecializationGetAll gets all specializations.
func (q *Queries) SpecializationGetAll(ctx context.Context, name string) ([]entity.Specialization, error) {
	name = "%%" + name + "%%"
	rows, err := q.db.Query(ctx, specializationGetAllSQL, name)

	if err != nil {
		return []entity.Specialization{}, errors.Wrap(err, "[queries.SpecializationGetAll] failed to get all specializations")
	}

	defer rows.Close()

	var specializations []entity.Specialization

	for rows.Next() {
		var specialization entity.Specialization

		err = rows.Scan(
			&specialization.ID,
			&specialization.Name,
		)

		if err != nil {
			return []entity.Specialization{}, errors.Wrap(err, "[queries.SpecializationGetAll] failed to scan specialization")
		}

		specializations = append(specializations, specialization)
	}

	return specializations, nil
}

// func specialization auto complete

func (q *Queries) SpecializationAutoComplete(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, specializationAutoCompleteSQL)

	if err != nil {
		return []string{}, errors.Wrap(err, "[queries.SpecializationAutoComplete] failed to get all specializations")
	}

	defer rows.Close()

	var specializations []string

	for rows.Next() {
		var specialization string

		err = rows.Scan(
			&specialization,
		)

		if err != nil {
			return []string{}, errors.Wrap(err, "[queries.SpecializationAutoComplete] failed to scan specialization")
		}

		specializations = append(specializations, specialization)
	}

	return specializations, nil
}
