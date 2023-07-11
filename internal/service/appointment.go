package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
	"github.com/pkg/errors"
)

// AppointmentCreate creates a new specialization Service.
func (s *Service) AppointmentCreate(ctx context.Context, arg entity.Appointment) (entity.Appointment, error) {
	var appointment entity.Appointment

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error
		arg.ID = uuid.New().String()

		appointment, err = q.AppointmentCreate(ctx, arg)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.Appointment{}, errors.Wrap(err, "[Service.AppointmentCreate] failed to create appointment")
	}

	return appointment, nil
}

// AppointmentServiceGetAll gets all appointment Service.
func (s *Service) AppointmentServiceGetAll(ctx context.Context) ([]entity.Appointment, error) {
	var appointments []entity.Appointment

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		appointments, err = q.AppointmentServiceGetAll(ctx)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[Service.AppointmentServiceGetAll] failed to get all appointments with Service")
	}

	return appointments, nil
}

// AppointmentDoctorGetAll gets all appointment Service.
func (s *Service) AppointmentDoctorGetAll(ctx context.Context) ([]entity.Appointment, error) {
	var appointments []entity.Appointment

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		appointments, err = q.AppointmentDoctorGetAll(ctx)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[Service.AppointmentDoctorGetAll] failed to get all appointments with doctor")
	}

	return appointments, nil
}

// AppointmentServiceGetAllOfPatient gets all appointments of patient Service.
func (s *Service) AppointmentServiceGetAllOfPatient(ctx context.Context, patientId string) ([]entity.Appointment, error) {
	var appointments []entity.Appointment

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		appointments, err = q.AppointmentServiceGetAllOfPatient(ctx, patientId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[Service.AppointmentServiceGetAllOfPatient] failed to get all appointments of patient with Service")
	}

	return appointments, nil
}

// AppointmentDoctorGetAllOfPatient gets all appointments of patient Service.
func (s *Service) AppointmentDoctorGetAllOfPatient(ctx context.Context, patientId string) ([]entity.Appointment, error) {
	var appointments []entity.Appointment

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		appointments, err = q.AppointmentDoctorGetAllOfPatient(ctx, patientId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[Service.AppointmentDoctorGetAllOfPatient] failed to get all appointments of patient with doctor")
	}

	return appointments, nil
}

// AppointmentGetAllOfDoctor gets all appointments of doctor Service.
func (s *Service) AppointmentGetAllOfDoctor(ctx context.Context, doctorId string) ([]entity.Appointment, error) {
	var appointments []entity.Appointment

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		appointments, err = q.AppointmentGetAllOfDoctor(ctx, doctorId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[Service.AppointmentGetAllOfDoctor] failed to get all appointments")
	}

	return appointments, nil
}

func (s *Service) AppointmentGetAllOfPatient(ctx context.Context, userId string) ([]entity.Appointment, error) {
	var appointments []entity.Appointment

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		appointments, err = q.AppointmentGetAllOfPatient(ctx, userId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[Service.AppointmentGetAllOfDoctor] failed to get all appointments")
	}

	return appointments, nil
}

// AppointmentGetAllOfService gets all appointments of Service Service.
func (s *Service) AppointmentGetAllOfService(ctx context.Context, serviceId string) ([]entity.Appointment, error) {
	var appointments []entity.Appointment

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		appointments, err = q.AppointmentGetAllOfService(ctx, serviceId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[Service.AppointmentGetAllOfService] failed to get all appointments")
	}

	return appointments, nil
}

// AppointmentApprove approves appointment Service.
func (s *Service) AppointmentApprove(ctx context.Context, id string) (entity.Appointment, error) {
	var appointment entity.Appointment

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		appointment, err = q.AppointmentApprove(ctx, true, id)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.Appointment{}, errors.Wrap(err, "[Service.AppointmentApprove] failed to approve appointment")
	}

	return appointment, nil
}
