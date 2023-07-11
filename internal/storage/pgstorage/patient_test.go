package pgstorage

import (
	"context"

	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/pkg/errors"
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
