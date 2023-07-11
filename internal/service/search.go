package service

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
	"github.com/pkg/errors"
)

// AutoComplete gets all doctors Service.
func (s *Service) AutoComplete(ctx context.Context) ([]string, error) {
	var result []string
	var doctors []string
	var services []string
	var specializations []string

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		doctors, err = q.DoctorAutoComplete(ctx)

		if err != nil {
			return err
		}

		services, err = q.ServiceAutoComplete(ctx)

		if err != nil {
			return err
		}

		specializations, err = q.SpecializationAutoComplete(ctx)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []string{}, errors.Wrap(err, "[Service.AutoComplete] failed to get all doctors")
	}

	result = append(doctors, services...)
	result = append(result, specializations...)

	return result, nil
}
