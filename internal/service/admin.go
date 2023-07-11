package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
	"github.com/pkg/errors"
)

// AdminCreate creates user with admin role Service.
func (s *Service) AdminCreate(ctx context.Context, arg entity.User) (entity.User, error) {
	var user entity.User

	password, errH := hashPassword(arg.Password)

	if errH != nil {
		return entity.User{}, errors.Wrap(errH, "[Service.AdminCreate] failed to hash password")
	}
	arg.ID = uuid.New().String()
	arg.Password = password

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		arg.Role = entity.RoleAdmin

		user, err = q.UserCreate(ctx, arg)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.User{}, errors.Wrap(err, "[Service.AdminCreate] failed to create admin")
	}

	return user, nil
}
