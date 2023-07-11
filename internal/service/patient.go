package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
	"github.com/pkg/errors"
)

// PatientCreate creates a new patient Service, first creates user and then patient
func (s *Service) PatientCreate(ctx context.Context, arg entity.Patient) (entity.Patient, error) {
	var patient entity.Patient

	password, errH := hashPassword(arg.User.Password)

	if errH != nil {
		return entity.Patient{}, errors.Wrap(errH, "[Service.PatientCreate] failed to hash password")
	}

	arg.User.Password = password

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		arg.User.ID = uuid.New().String()
		arg.User.Role = entity.RolePatient

		user, err := q.UserCreate(ctx, arg.User)

		if err != nil {
			return err
		}

		arg.ID = uuid.New().String()

		patient, err = q.PatientCreate(ctx, arg)

		if err != nil {
			return err
		}

		patient.User = user

		return nil
	})

	if err != nil {
		return entity.Patient{}, errors.Wrap(err, "[Service.PatientCreate] failed to create patient")
	}

	return patient, nil
}

// PatientUpdate updates a patient Service.
func (s *Service) PatientUpdate(ctx context.Context, userId string, arg entity.PatientUpdateInput) (entity.Patient, error) {
	var patient entity.Patient

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		patient, err = q.PatientUpdate(ctx, userId, arg)

		if err != nil {
			return err
		}

		user, err := q.UserGetByID(ctx, userId)

		if err != nil {
			return err
		}

		patient.User = user

		return nil
	})

	if err != nil {
		return entity.Patient{}, errors.Wrap(err, "[Service.PatientUpdate] failed to update patient")
	}

	return patient, nil
}

// PatientGetAll gets all patients Service.
func (s *Service) PatientGetAll(ctx context.Context) ([]entity.Patient, error) {
	var patients []entity.Patient

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		patients, err = q.PatientGetAll(ctx)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Patient{}, errors.Wrap(err, "[Service.PatientGetAll] failed to get all patients")
	}

	return patients, nil
}

// PatientGetByUserID gets a patient by user id Service.
func (s *Service) PatientGetByUserID(ctx context.Context, userId string) (entity.Patient, error) {
	var patient entity.Patient

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		patient, err = q.PatientGetByUserID(ctx, userId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.Patient{}, errors.Wrap(err, "[Service.PatientGetByUserID] failed to get patient by user id")
	}

	return patient, nil
}
