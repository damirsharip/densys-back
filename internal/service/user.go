package service

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
	"github.com/pkg/errors"
)

// UserUpdate updates a user Service.
func (s *Service) UserUpdate(ctx context.Context, userId string, arg entity.UserUpdateInput) (entity.User, error) {
	var user entity.User

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		user, err = q.UserUpdate(ctx, userId, arg)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.User{}, errors.Wrap(err, "[Service.UserUpdate] failed to update user")
	}

	return user, nil
}

// UserGetByEmail returns a single user by email.
func (s *Service) UserGetByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		user, err = q.UserGetByEmail(ctx, email)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.User{}, errors.Wrap(err, "[Service.UserGetByEmail] failed to get user")
	}

	return user, nil
}
