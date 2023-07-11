package service

import (
	"context"
	"sort"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
)

// DoctorCreate creates a new doctor Service, first creates user and then doctor
func (s *Service) DoctorCreate(ctx context.Context, arg entity.Doctor) (entity.Doctor, error) {
	var doctor entity.Doctor

	password, errH := hashPassword(arg.User.Password)

	if errH != nil {
		return entity.Doctor{}, errors.Wrap(errH, "[Service.PatientCreate] failed to hash password")
	}

	if arg.Photo != "" {
		var errI error
		arg.Photo, errI = s.imageUpload(ctx, &arg.Photo)

		if errI != nil {
			return entity.Doctor{}, errors.Wrap(errI, "[Service.DoctorCreate] failed to upload photo")
		}
	}

	arg.User.Password = password

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		arg.User.ID = uuid.New().String()
		arg.User.Role = entity.RoleDoctor

		user, err := q.UserCreate(ctx, arg.User)

		if err != nil {
			return err
		}

		arg.ID = uuid.New().String()

		doctor, err = q.DoctorCreate(ctx, arg)

		if err != nil {
			return err
		}

		doctor.User = user

		return nil
	})

	if err != nil {
		return entity.Doctor{}, errors.Wrap(err, "[Service.DoctorCreate] failed to create doctor")
	}

	return doctor, nil
}

// DoctorUpdate updates a doctor Service.
func (s *Service) DoctorUpdate(ctx context.Context, userId string, arg entity.DoctorUpdateInput) (entity.Doctor, error) {
	var doctor entity.Doctor

	if arg.Photo != nil && *arg.Photo != "" {
		var errI error
		*arg.Photo, errI = s.imageUpload(ctx, arg.Photo)

		if errI != nil {
			return entity.Doctor{}, errors.Wrap(errI, "[Service.DoctorCreate] failed to upload photo")
		}
	}

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		doctor, err = q.DoctorUpdate(ctx, userId, arg)

		if err != nil {
			return err
		}

		user, err := q.UserGetByID(ctx, userId)

		if err != nil {
			return err
		}

		doctor.User = user

		return nil
	})

	if err != nil {
		return entity.Doctor{}, errors.Wrap(err, "[Service.DoctorUpdate] failed to update doctor")
	}

	return doctor, nil
}

// DoctorGetAll gets all doctors Service.
func (s *Service) DoctorGetAll(ctx context.Context, name string) ([]entity.Doctor, error) {
	var doctors []entity.Doctor

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		doctors, err = q.DoctorGetAll(ctx, name)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Doctor{}, errors.Wrap(err, "[Service.DoctorGetAll] failed to get all doctors")
	}

	return doctors, nil
}

// DoctorSearchByServiceAndSpecialization gets doctors by Service and specialization Service.
func (s *Service) DoctorSearchByServiceAndSpecialization(ctx context.Context, search string, specializationId string, pagination entity.Pagination) ([]entity.Doctor, entity.Pagination, error) {
	var doctors []entity.Doctor

	pagination.PerPage = 6

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		doctors, err = q.DoctorSearchByServiceAndSpecialization(ctx, search, specializationId, pagination)

		if err != nil {
			return err
		}

		for i := range doctors {
			doctors[i].Services, _ = q.DoctorServiceGetAll(ctx, doctors[i].User.ID)
		}

		pagination.TotalCount, err = q.DoctorSearchByServiceAndSpecializationCount(ctx, search, specializationId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return []entity.Doctor{}, entity.Pagination{}, errors.Wrap(err, "[Service.DoctorGetAll] failed to get all doctors")
	}

	pagination.TotalPages = pagination.TotalCount/pagination.PerPage + 1

	return doctors, pagination, nil
}

// DoctorGetByUserID gets a doctor by user id Service.
func (s *Service) DoctorGetByUserID(ctx context.Context, userId string) (entity.Doctor, error) {
	var doctor entity.Doctor

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		doctor, err = q.DoctorGetByUserID(ctx, userId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.Doctor{}, errors.Wrap(err, "[Service.DoctorGetByUserID] failed to get doctor by user id")
	}

	return doctor, nil
}

func (s *Service) DoctorGetAllBySpecialization(ctx context.Context, id string, date string, pageNumber int) ([]entity.DoctorWithSchedule, error) {
	var doctors []entity.DoctorWithSchedule
	var appointments []entity.Appointment

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		doctors, err = q.DoctorGetAllBySpecialization(ctx, id, (pageNumber-1)*10)

		if err != nil {
			return err
		}

		for i, v := range doctors {
			appointments, err = q.AppointmentGetAllByDoctorAndDate(ctx, v.User.ID, date)

			if err != nil {
				return err
			}

			m := map[string]int{
				"09:00": 1, "10:00": 1, "11:00": 1,
				"12:00": 1, "14:00": 1, "15:00": 1,
				"16:00": 1, "17:00": 1, "18:00": 1,
			}

			for _, d := range appointments {
				time := d.Time.String()[11:16]
				if _, ok := m[time]; ok { // if exists make 0
					m[time] = 0
				}
			}

			for j, d := range m { // traverse over map
				if d == 1 { // and take the ones that does not exist in appointment times
					doctors[i].Schedule = append(doctors[i].Schedule, j)
				}
			}

			sort.Strings(doctors[i].Schedule) // sort them to look fine

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return []entity.DoctorWithSchedule{}, errors.Wrap(err, "[Service.DoctorGetAllBySpecialization] failed to get all doctors by specialization")
	}

	return doctors, nil
}

func (s *Service) DoctorGetScheduleByDate(ctx context.Context, userId, date string) ([]entity.Schedule, error) {
	var schedule []entity.Schedule
	var appointments []entity.Appointment

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(q *pgstorage.Queries) error {
		var err error

		appointments, err = q.AppointmentGetAllByDoctorAndDate(ctx, userId, date)

		if err != nil {
			return err
		}

		return nil
	})

	m := map[string]int{
		"09:00": 1, "10:00": 1, "11:00": 1,
		"12:00": 1, "14:00": 1, "15:00": 1,
		"16:00": 1, "17:00": 1, "18:00": 1,
	}

	for _, d := range appointments {
		time := d.Time.String()[11:16]
		if _, ok := m[time]; ok { // if exists make 0
			m[time] = 0
		}
	}

	for j, d := range m { // traverse over map
		if d == 1 { // and take the ones that does not exist in appointment times
			schedule = append(schedule, entity.Schedule(j))
		}
	}

	sort.Slice(schedule, func(i, j int) bool {
		return schedule[i] < schedule[j]
	})

	if err != nil {
		return []entity.Schedule{}, errors.Wrap(err, "[Service.DoctorGetAllBySpecialization] failed to get all doctors by specialization")
	}

	return schedule, nil
}
