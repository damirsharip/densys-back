package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
	"github.com/pkg/errors"
)

// SpecializationCreate creates a new specialization Service.
func (s *Service) SpecializationCreate(ctx context.Context, arg entity.Specialization) (entity.Specialization, error) {
	var specialization entity.Specialization

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		arg.ID = uuid.New().String()

		specialization, err = q.SpecializationCreate(ctx, arg)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.Specialization{}, errors.Wrap(err, "[Service.SpecializationCreate] failed to create specialization")
	}

	return specialization, nil
}

// SpecializationGetAll gets all specializations Service.
func (s *Service) SpecializationGetAll(ctx context.Context, name string) ([]entity.Specialization, error) {
	var specializations []entity.Specialization

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		specializations, err = q.SpecializationGetAll(ctx, name)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "[Service.SpecializationGetAll] failed to get all specializations")
	}

	return specializations, nil
}
