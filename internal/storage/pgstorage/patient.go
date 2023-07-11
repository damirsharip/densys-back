package pgstorage

import (
	"context"

	"github.com/pkg/errors"

	"github.com/khanfromasia/densys/admin/internal/entity"
)

// PatientCreate creates a new patient.
func (q *Queries) PatientCreate(ctx context.Context, arg entity.Patient) (entity.Patient, error) {
	_, err := q.db.Exec(
		ctx,
		patientCreateSQL,
		arg.ID,
		arg.User.ID,
		arg.BloodGroup,
		arg.EmergencyContactNumber,
		arg.MaritalStatus,
	)

	if err != nil {
		return entity.Patient{}, errors.Wrap(err, "[queries.PatientCreate] failed to create patient")
	}

	return arg, nil
}

// PatientGetByUserID gets a patient by user ID.
func (q *Queries) PatientGetByUserID(ctx context.Context, userID string) (entity.Patient, error) {
	var patient entity.Patient

	err := q.db.QueryRow(
		ctx,
		patientGetByUserIDSQL,
		userID,
	).Scan(
		&patient.ID,
		&patient.User.ID,
		&patient.User.FirstName,
		&patient.User.LastName,
		&patient.User.MiddleName,
		&patient.User.Email,
		&patient.User.PhoneNumber,
		&patient.User.Address,
		&patient.User.BirthDate,
		&patient.User.IIN,
		&patient.User.GovernmentID,
		&patient.User.Password,
		&patient.User.Role,
		&patient.User.CreatedAt,
		&patient.BloodGroup,
		&patient.EmergencyContactNumber,
		&patient.MaritalStatus,
	)

	if err != nil {
		return entity.Patient{}, errors.Wrap(err, "[queries.PatientGetByUserID] failed to get patient by user ID")
	}

	return patient, nil
}

// PatientGetAll gets all patients.
func (q *Queries) PatientGetAll(ctx context.Context) ([]entity.Patient, error) {
	rows, err := q.db.Query(ctx, patientGetAllSQL)

	if err != nil {
		return []entity.Patient{}, errors.Wrap(err, "[queries.PatientGetAll] failed to get all patients")
	}

	defer rows.Close()

	var patients []entity.Patient

	for rows.Next() {
		var patient entity.Patient

		err = rows.Scan(
			&patient.ID,
			&patient.User.ID,
			&patient.User.FirstName,
			&patient.User.LastName,
			&patient.User.Email,
			&patient.User.PhoneNumber,
			&patient.User.Password,
			&patient.User.Role,
			&patient.User.CreatedAt,
			&patient.User.BirthDate,
			&patient.BloodGroup,
			&patient.EmergencyContactNumber,
			&patient.MaritalStatus,
		)

		if err != nil {
			return []entity.Patient{}, errors.Wrap(err, "[queries.PatientGetAll] failed to scan patient")
		}

		patients = append(patients, patient)
	}

	return patients, nil
}

// PatientUpdate updates a patient.
func (q *Queries) PatientUpdate(ctx context.Context, userId string, arg entity.PatientUpdateInput) (entity.Patient, error) {
	row := q.db.QueryRow(
		ctx,
		patientUpdateSQL,
		arg.BloodGroup,
		arg.EmergencyContactNumber,
		arg.MaritalStatus,
		userId,
	)

	var patient entity.Patient

	err := row.Scan(
		&patient.ID,
		&patient.User.ID,
		&patient.BloodGroup,
		&patient.EmergencyContactNumber,
		&patient.MaritalStatus,
	)

	if err != nil {
		return entity.Patient{}, errors.Wrap(err, "[queries.PatientUpdate] failed to update patient")
	}

	return patient, nil
}
