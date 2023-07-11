package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
	"github.com/pkg/errors"
)

// ServiceCreate creates a new department Service.
func (s *Service) ServiceCreate(ctx context.Context, arg entity.Service) (entity.Service, error) {
	var procedure entity.Service

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		arg.ID = uuid.New().String()

		procedure, err = q.ServiceCreate(ctx, arg)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.Service{}, errors.Wrap(err, "[Service.DepartmentCreate] failed to create department")
	}

	return procedure, nil
}

// ServiceGetAll gets all departments Service.
func (s *Service) ServiceGetAll(ctx context.Context) ([]entity.Service, error) {
	var procedures []entity.Service

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		procedures, err = q.ServiceGetAll(ctx)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Service{}, errors.Wrap(err, "[Service.DepartmentGetAll] failed to get all departments")
	}

	return procedures, nil
}

// DoctorServiceGetAll gets all departments Service.
func (s *Service) DoctorServiceGetAll(ctx context.Context, userId string) ([]entity.Service, error) {
	var procedures []entity.Service

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		procedures, err = q.DoctorServiceGetAll(ctx, userId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Service{}, errors.Wrap(err, "[Service.DepartmentGetAll] failed to get all departments")
	}

	return procedures, nil
}
