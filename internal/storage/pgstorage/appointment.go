package pgstorage

import (
	"context"

	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/pkg/errors"
)

// AppointmentCreate creates a new appointment.
func (q *Queries) AppointmentCreate(ctx context.Context, arg entity.Appointment) (entity.Appointment, error) {
	args := []interface{}{
		arg.ID,
		&arg.Patient.ID,
		&arg.Doctor.ID,
		&arg.Service.ID,
		arg.Date,
		arg.Time,
	}

	if arg.Service.ID == "" {
		args[3] = nil
	} else if arg.Doctor.ID == "" {
		args[2] = nil
	}

	_, err := q.db.Exec(
		ctx,
		appointmentCreateSQL,
		args...,
	)

	if err != nil {
		return entity.Appointment{}, err
	}

	return arg, nil
}

// AppointmentServiceGetAll gets all appointments of a patient.
func (q *Queries) AppointmentServiceGetAll(ctx context.Context) ([]entity.Appointment, error) {
	rows, err := q.db.Query(ctx, appointmentServiceGetAllSQL)

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentServiceGetAll] failed to get all appointments with service")
	}

	defer rows.Close()

	var appointments []entity.Appointment

	for rows.Next() {
		var appointment entity.Appointment

		err = rows.Scan(
			&appointment.ID,
			&appointment.Date,
			&appointment.Time,
			&appointment.IsApproved,
			&appointment.Patient.ID,
			&appointment.Patient.FirstName,
			&appointment.Patient.LastName,
			&appointment.Service.ID,
			&appointment.Service.Name,
			&appointment.Service.Price,
			&appointment.Service.ListOfContradictions,
			&appointment.Service.OtherRelatedInformation,
		)

		if err != nil {
			return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentServiceGetAll] failed to scan appointment with service")
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// AppointmentDoctorGetAll gets all appointments of a patient.
func (q *Queries) AppointmentDoctorGetAll(ctx context.Context) ([]entity.Appointment, error) {
	rows, err := q.db.Query(ctx, appointmentDoctorGetAllSQL)

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentDoctorGetAll] failed to get all appointments with doctor")
	}

	defer rows.Close()

	var appointments []entity.Appointment

	for rows.Next() {
		var appointment entity.Appointment

		err = rows.Scan(
			&appointment.ID,
			&appointment.Date,
			&appointment.Time,
			&appointment.IsApproved,
			&appointment.Patient.ID,
			&appointment.Patient.FirstName,
			&appointment.Patient.LastName,
			&appointment.Doctor.ID,
			&appointment.Doctor.FirstName,
			&appointment.Doctor.LastName,
			&appointment.Service.ID,
			&appointment.Service.Name,
			&appointment.Service.Price,
			&appointment.Service.ListOfContradictions,
			&appointment.Service.OtherRelatedInformation,
		)

		if err != nil {
			return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentDoctorGetAll] failed to scan appointment with doctor")
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// AppointmentServiceGetAllOfPatient gets all appointments of a patient.
func (q *Queries) AppointmentServiceGetAllOfPatient(ctx context.Context, patientId string) ([]entity.Appointment, error) {
	rows, err := q.db.Query(ctx, appointmentServiceGetAllOfPatientSQL, patientId)

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentServiceGetAllOfPatient] failed to get all appointments of patient with service")
	}

	defer rows.Close()

	var appointments []entity.Appointment

	for rows.Next() {
		var appointment entity.Appointment

		err = rows.Scan(
			&appointment.ID,
			&appointment.Date,
			&appointment.Time,
			&appointment.IsApproved,
			&appointment.Service.ID,
			&appointment.Service.Name,
			&appointment.Service.Price,
			&appointment.Service.ListOfContradictions,
			&appointment.Service.OtherRelatedInformation,
		)

		if err != nil {
			return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentServiceGetAllOfPatient] failed to scan appointment of patient with service")
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// AppointmentDoctorGetAllOfPatient gets all appointments of a patient.
func (q *Queries) AppointmentDoctorGetAllOfPatient(ctx context.Context, patientId string) ([]entity.Appointment, error) {
	rows, err := q.db.Query(ctx, appointmentDoctorGetAllOfPatientSQL, patientId)

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentDoctorGetAllOfPatient] failed to get all appointments of patient with doctor")
	}

	defer rows.Close()

	var appointments []entity.Appointment

	for rows.Next() {
		var appointment entity.Appointment

		err = rows.Scan(
			&appointment.ID,
			&appointment.Date,
			&appointment.Time,
			&appointment.IsApproved,
			&appointment.Patient.ID,
			&appointment.Patient.FirstName,
			&appointment.Patient.LastName,
			&appointment.Doctor.ID,
			&appointment.Doctor.FirstName,
			&appointment.Doctor.LastName,
			&appointment.Service.ID,
			&appointment.Service.Name,
			&appointment.Service.Price,
			&appointment.Service.ListOfContradictions,
			&appointment.Service.OtherRelatedInformation,
		)

		if err != nil {
			return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentDoctorGetAllOfPatient] failed to scan appointment of patient with doctor")
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// AppointmentGetAllOfDoctor gets all specializations.
func (q *Queries) AppointmentGetAllOfDoctor(ctx context.Context, doctorId string) ([]entity.Appointment, error) {
	rows, err := q.db.Query(ctx, appointmentGetAllOfDoctorSQL, doctorId)

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentGetAllOfDoctor] failed to get all specializations")
	}

	defer rows.Close()

	var appointments []entity.Appointment

	for rows.Next() {
		var appointment entity.Appointment

		err = rows.Scan(
			&appointment.ID,
			&appointment.Date,
			&appointment.Time,
			&appointment.IsApproved,
			&appointment.Patient.ID,
			&appointment.Patient.FirstName,
			&appointment.Patient.LastName,
			&appointment.Doctor.ID,
			&appointment.Doctor.FirstName,
			&appointment.Doctor.LastName,
			&appointment.Service.ID,
			&appointment.Service.Name,
			&appointment.Service.Price,
			&appointment.Service.ListOfContradictions,
			&appointment.Service.OtherRelatedInformation,
		)

		if err != nil {
			return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentGetAllOfDoctor] failed to scan specialization")
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// AppointmentGetAllOfPatient gets all specializations.
func (q *Queries) AppointmentGetAllOfPatient(ctx context.Context, userId string) ([]entity.Appointment, error) {
	rows, err := q.db.Query(ctx, appointmentGetAllOfPatientSQL, userId)

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentGetAllOfPatient] failed to get all specializations")
	}

	defer rows.Close()

	var appointments []entity.Appointment

	for rows.Next() {
		var appointment entity.Appointment

		err = rows.Scan(
			&appointment.ID,
			&appointment.Date,
			&appointment.Time,
			&appointment.IsApproved,
			&appointment.Patient.ID,
			&appointment.Patient.FirstName,
			&appointment.Patient.LastName,
			&appointment.Doctor.ID,
			&appointment.Doctor.FirstName,
			&appointment.Doctor.LastName,
			&appointment.Service.ID,
			&appointment.Service.Name,
			&appointment.Service.Price,
			&appointment.Service.ListOfContradictions,
			&appointment.Service.OtherRelatedInformation,
		)

		if err != nil {
			return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentGetAllOfPatient] failed to scan specialization")
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// AppointmentGetAllOfService gets all specializations.
func (q *Queries) AppointmentGetAllOfService(ctx context.Context, serviceId string) ([]entity.Appointment, error) {
	rows, err := q.db.Query(ctx, appointmentGetAllOfServiceSQL, serviceId)

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentGetAllOfService] failed to get all specializations")
	}

	defer rows.Close()

	var appointments []entity.Appointment

	for rows.Next() {
		var appointment entity.Appointment

		err = rows.Scan(
			&appointment.ID,
			&appointment.Date,
			&appointment.Time,
			&appointment.IsApproved,
			&appointment.Patient.ID,
			&appointment.Patient.FirstName,
			&appointment.Patient.LastName,
		)

		if err != nil {
			return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentGetAllOfService] failed to scan specialization")
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// AppointmentApprove gets all appointments of a patient.
func (q *Queries) AppointmentApprove(ctx context.Context, isApproved bool, id string) (entity.Appointment, error) {
	row := q.db.QueryRow(ctx, appointmentApproveSQL, id, isApproved)

	var appointment entity.Appointment

	if err := row.Scan(
		&appointment.ID,
		&appointment.IsApproved,
	); err != nil {
		return entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentApprove] failed to update appointment")
	}

	return appointment, nil
}
