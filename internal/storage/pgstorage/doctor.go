package pgstorage

import (
	"context"

	"github.com/pkg/errors"

	"github.com/khanfromasia/densys/admin/internal/entity"
)

// DoctorCreate creates a new doctor.
func (q *Queries) DoctorCreate(ctx context.Context, arg entity.Doctor) (entity.Doctor, error) {
	args := []interface{}{
		arg.ID,
		arg.User.ID,
		arg.Department.ID,
		arg.Specialization.ID,
		arg.WorkExperience,
		arg.Photo,
		arg.Category,
		arg.ScheduleDetails,
		arg.Degree,
		arg.AppointmentPrice,
		arg.Rating,
	}

	_, err := q.db.Exec(
		ctx,
		doctorCreateSQL,
		args...,
	)

	if err != nil {
		return entity.Doctor{}, err
	}

	return arg, nil
}

// DoctorGetByUserID joined user, department, and specialization.
func (q *Queries) DoctorGetByUserID(ctx context.Context, id string) (entity.Doctor, error) {
	var doctor entity.Doctor

	err := q.db.QueryRow(
		ctx,
		doctorGetByUserIDSQL,
		id,
	).Scan(
		&doctor.ID,
		&doctor.User.ID,
		&doctor.User.FirstName,
		&doctor.User.LastName,
		&doctor.User.MiddleName,
		&doctor.User.Email,
		&doctor.User.PhoneNumber,
		&doctor.User.Address,
		&doctor.User.BirthDate,
		&doctor.User.IIN,
		&doctor.User.GovernmentID,
		&doctor.User.Password,
		&doctor.User.Role,
		&doctor.User.CreatedAt,
		&doctor.Department.ID,
		&doctor.Department.Name,
		&doctor.Specialization.ID,
		&doctor.Specialization.Name,
		&doctor.WorkExperience,
		&doctor.Photo,
		&doctor.Category,
		&doctor.ScheduleDetails,
		&doctor.Degree,
		&doctor.AppointmentPrice,
		&doctor.Rating,
	)

	if err != nil {
		return entity.Doctor{}, errors.Wrap(err, "[queries.DoctorGetByID] failed to get doctor by id")
	}

	return doctor, nil
}

func (q *Queries) DoctorSearchByServiceAndSpecializationCount(ctx context.Context, search string, specializationId string) (int64, error) {
	var count int64
	search = "%%" + search + "%%"
	var sid *string
	if specializationId != "" {
		sid = &specializationId
	}

	err := q.db.QueryRow(ctx, doctorFilterSearchBySpecializationAndServiceCount, search, sid).Scan(&count)

	if err != nil {
		return 0, errors.Wrap(err, "[queries.DoctorSearchByServiceAndSpecializationCount] failed to get count")
	}

	return count, nil
}

// DoctorSearchByServiceAndSpecialization doctorFilterSearchBySpecializationAndService
func (q *Queries) DoctorSearchByServiceAndSpecialization(ctx context.Context, search string, specializationId string, pagination entity.Pagination) ([]entity.Doctor, error) {
	search = "%%" + search + "%%"
	var sid *string
	if specializationId != "" {
		sid = &specializationId
	}
	rows, err := q.db.Query(ctx, doctorFilterSearchBySpecializationAndService, search, sid, pagination.PerPage, (pagination.CurrentPage-1)*pagination.PerPage)

	if err != nil {
		return []entity.Doctor{}, errors.Wrap(err, "[queries.DoctorSearchByServiceAndSpecialization] failed to get all doctors")
	}

	defer rows.Close()

	var doctors []entity.Doctor

	for rows.Next() {
		var doctor entity.Doctor

		err = rows.Scan(
			&doctor.ID,
			&doctor.User.ID,
			&doctor.User.FirstName,
			&doctor.User.LastName,
			&doctor.User.Email,
			&doctor.User.PhoneNumber,
			&doctor.User.Password,
			&doctor.User.Role,
			&doctor.User.CreatedAt,
			&doctor.User.BirthDate,
			&doctor.Department.ID,
			&doctor.Department.Name,
			&doctor.Specialization.ID,
			&doctor.Specialization.Name,
			&doctor.WorkExperience,
			&doctor.Photo,
			&doctor.Category,
			&doctor.ScheduleDetails,
			&doctor.Degree,
			&doctor.AppointmentPrice,
			&doctor.Rating,
		)

		if err != nil {
			return []entity.Doctor{}, errors.Wrap(err, "[queries.DoctorSearchByServiceAndSpecialization] failed to scan doctor")
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

// DoctorGetAll joined user, department, and specialization.
func (q *Queries) DoctorGetAll(ctx context.Context, name string) ([]entity.Doctor, error) {
	name = "%%" + name + "%%"
	rows, err := q.db.Query(ctx, doctorGetAllSQL, name)

	if err != nil {
		return []entity.Doctor{}, errors.Wrap(err, "[queries.DoctorGetAll] failed to get all doctors")
	}

	defer rows.Close()

	var doctors []entity.Doctor

	for rows.Next() {
		var doctor entity.Doctor

		err = rows.Scan(
			&doctor.ID,
			&doctor.User.ID,
			&doctor.User.FirstName,
			&doctor.User.LastName,
			&doctor.User.Email,
			&doctor.User.PhoneNumber,
			&doctor.User.Password,
			&doctor.User.Role,
			&doctor.User.CreatedAt,
			&doctor.User.BirthDate,
			&doctor.Department.ID,
			&doctor.Department.Name,
			&doctor.Specialization.ID,
			&doctor.Specialization.Name,
			&doctor.WorkExperience,
			&doctor.Photo,
			&doctor.Category,
			&doctor.ScheduleDetails,
			&doctor.Degree,
			&doctor.AppointmentPrice,
			&doctor.Rating,
		)

		if err != nil {
			return []entity.Doctor{}, errors.Wrap(err, "[queries.DoctorGetAll] failed to scan doctor")
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

// DoctorAutocomplete joined user, department, and specialization.
func (q *Queries) DoctorAutoComplete(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, doctorAutoCompleteSQL)

	if err != nil {
		return []string{}, errors.Wrap(err, "[queries.DoctorAutoComplete] failed to get all doctors")
	}

	defer rows.Close()

	var doctors []string

	for rows.Next() {
		var doctor string

		err = rows.Scan(&doctor)

		if err != nil {
			return []string{}, errors.Wrap(err, "[queries.DoctorAutoComplete] failed to scan doctor")
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

// DoctorUpdate updates a doctor.
func (q *Queries) DoctorUpdate(ctx context.Context, userId string, arg entity.DoctorUpdateInput) (entity.Doctor, error) {
	args := []interface{}{
		arg.DepartmentID,
		arg.SpecializationID,
		arg.WorkExperience,
		arg.Photo,
		arg.Category,
		arg.ScheduleDetails,
		arg.Degree,
		arg.AppointmentPrice,
		arg.Rating,
		userId,
	}

	row := q.db.QueryRow(
		ctx,
		doctorUpdateSQL,
		args...,
	)

	var doctor entity.Doctor

	if err := row.Scan(
		&doctor.ID,
		&doctor.User.ID,
		&doctor.Department.ID,
		&doctor.Specialization.ID,
		&doctor.WorkExperience,
		&doctor.Photo,
		&doctor.Category,
		&doctor.ScheduleDetails,
		&doctor.Degree,
		&doctor.AppointmentPrice,
		&doctor.Rating,
	); err != nil {
		return entity.Doctor{}, errors.Wrap(err, "[queries.DoctorUpdate] failed to update doctor")
	}

	return doctor, nil
}

// DoctorGetAllBySpecialization joined user, department, and specialization and selects with particular specialization id.
func (q *Queries) DoctorGetAllBySpecialization(ctx context.Context, id string, offset int) ([]entity.DoctorWithSchedule, error) {
	rows, err := q.db.Query(ctx, doctorGetAllBySpecializationSQL, id, offset)

	if err != nil {
		return []entity.DoctorWithSchedule{}, errors.Wrap(err, "[queries.DoctorGetAllBySpecialization] failed to get all doctors by specialization")
	}

	defer rows.Close()

	var doctors []entity.DoctorWithSchedule

	for rows.Next() {
		var doctor entity.DoctorWithSchedule

		err = rows.Scan(
			&doctor.ID,
			&doctor.User.ID,
			&doctor.User.FirstName,
			&doctor.User.LastName,
			&doctor.User.Email,
			&doctor.User.PhoneNumber,
			&doctor.User.Password,
			&doctor.User.Role,
			&doctor.User.CreatedAt,
			&doctor.User.BirthDate,
			&doctor.Department.ID,
			&doctor.Department.Name,
			&doctor.Specialization.ID,
			&doctor.Specialization.Name,
			&doctor.WorkExperience,
			&doctor.Photo,
			&doctor.Category,
			&doctor.ScheduleDetails,
			&doctor.Degree,
			&doctor.AppointmentPrice,
			&doctor.Rating,
		)
		if err != nil {
			return []entity.DoctorWithSchedule{}, errors.Wrap(err, "[queries.DoctorGetAllBySpecialization] failed to scan doctor")
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (q *Queries) AppointmentGetAllByDoctorAndDate(ctx context.Context, doctorId, date string) ([]entity.Appointment, error) {
	rows, err := q.db.Query(ctx, appointmentGetAllByDoctorAndDateSQL, doctorId, date)

	if err != nil {
		return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentGetAllByDoctorAndDate] failed to get all doctors by specialization")
	}

	defer rows.Close()

	var appointments []entity.Appointment

	for rows.Next() {
		var appointment entity.Appointment

		err = rows.Scan(
			&appointment.Date,
			&appointment.Time,
		)
		if err != nil {
			return []entity.Appointment{}, errors.Wrap(err, "[queries.AppointmentGetAllByDoctorAndDate] failed to scan doctor")
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}
