package pgstorage

import (
	"context"

	"github.com/pkg/errors"

	"github.com/khanfromasia/densys/admin/internal/entity"
)

// UserCreate creates a new user.
func (q *Queries) UserCreate(ctx context.Context, arg entity.User) (entity.User, error) {
	_, err := q.db.Exec(
		ctx,
		userCreateSQL,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.MiddleName,
		arg.Address,
		arg.PhoneNumber,
		arg.Email,
		arg.Role,
		arg.Password,
		arg.BirthDate,
		arg.IIN,
		arg.GovernmentID,
		arg.CreatedAt,
	)

	if err != nil {
		return entity.User{}, errors.Wrap(err, "[queries.UserCreate] error while inserting user")
	}

	// if !ct.Insert() || ct.RowsAffected() != 1 {}

	return arg, nil
}

// UserGetByID returns a single user by ID.
func (q *Queries) UserGetByID(ctx context.Context, id string) (entity.User, error) {
	row := q.db.QueryRow(ctx, userGetByIDSQL, id)

	var user entity.User
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Address,
		&user.PhoneNumber,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.BirthDate,
		&user.IIN,
		&user.GovernmentID,
		&user.CreatedAt,
	); err != nil {
		return entity.User{}, errors.Wrap(err, "[queries.UserGetByID] error while getting user")
	}

	return user, nil
}

// UserUpdate updates a user.
func (q *Queries) UserUpdate(ctx context.Context, userId string, arg entity.UserUpdateInput) (entity.User, error) {
	row := q.db.QueryRow(
		ctx,
		userUpdateSQL,
		arg.FirstName,
		arg.LastName,
		arg.MiddleName,
		arg.Address,
		arg.PhoneNumber,
		arg.Email,
		arg.BirthDate,
		arg.IIN,
		arg.GovernmentID,
		userId,
	)

	var user entity.User
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Address,
		&user.PhoneNumber,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.BirthDate,
		&user.IIN,
		&user.GovernmentID,
		&user.CreatedAt,
	); err != nil {
		return entity.User{}, errors.Wrap(err, "[queries.UserUpdate] error while updating user")
	}

	return user, nil
}

// UserGetByEmail returns a single user by email.
func (q *Queries) UserGetByEmail(ctx context.Context, email string) (entity.User, error) {
	row := q.db.QueryRow(ctx, userGetByEmailSQL, email)

	var user entity.User
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Address,
		&user.PhoneNumber,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.BirthDate,
		&user.IIN,
		&user.GovernmentID,
		&user.CreatedAt,
	); err != nil {
		return entity.User{}, errors.Wrap(err, "[queries.UserGetByEmail] error while getting user")
	}

	return user, nil
}
