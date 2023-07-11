package pgstorage

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage/fixture"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage_UserCreate(t *testing.T) {
	// arrange
	f := setUp(t)
	defer f.tearDown()

	// act
	v := fixture.User().
		ID(uuid.New().String()).
		FirstName("John").
		LastName("Doe").
		MiddleName("Joe").
		Address("123 Main St").
		Email("john.doe@gmail.com").
		Password("password").
		BirthDate(time.Now()).
		Role("patient").
		GovernmentID("1234567890").
		IIN("123456789012").
		PhoneNumber("1234567890").
		CreatedAt(time.Now()).
		V()

	querySelect := regexp.QuoteMeta(
		`INSERT INTO public.users 
			(id,first_name,last_name,middle_name,address,phone_number,email,role,password,birth_date,iin,government_id,created_at)
		VALUES 
			($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`,
	)

	args := []interface{}{
		v.ID,
		v.FirstName,
		v.LastName,
		v.MiddleName,
		v.Address,
		v.Email,
		v.Role,
		v.Password,
		v.BirthDate,
		v.IIN,
		v.GovernmentID,
		v.CreatedAt,
	}

	res := pgxmock.NewResult("INSERT", 1)

	f.pgxPoolMock.ExpectExec(querySelect).
		WithArgs(args...).
		WillReturnResult(res)

	ctx := context.Background()

	user, err := f.storage.UserCreate(ctx, v)

	expectErr := f.pgxPoolMock.ExpectationsWereMet()

	// assert
	require.NoError(t, err)
	require.NoError(t, expectErr)

	require.NotEmpty(t, user)

	assert.Equal(t, v.ID, user.ID)
	assert.Equal(t, v.FirstName, user.FirstName)
	assert.Equal(t, v.LastName, user.LastName)
	assert.Equal(t, v.MiddleName, user.MiddleName)
	assert.Equal(t, v.PhoneNumber, user.PhoneNumber)
	assert.Equal(t, v.IIN, user.IIN)
	assert.Equal(t, v.GovernmentID, user.GovernmentID)
	assert.Equal(t, v.Address, user.Address)
	assert.Equal(t, v.Email, user.Email)
	assert.Equal(t, v.Role, user.Role)
	assert.Equal(t, v.BirthDate, user.BirthDate)
	assert.Equal(t, v.CreatedAt, user.CreatedAt)
}

func TestStorage_UserGetByID(t *testing.T) {
	// arrange
	f := setUp(t)
	defer f.tearDown()

	// act
	v := fixture.User().
		ID(uuid.New().String()).
		FirstName("John").
		LastName("Doe").
		MiddleName("Joe").
		Address("123 Main St").
		Email("john.doe@gmail.com").
		Password("password").
		BirthDate(time.Now()).
		Role("patient").
		GovernmentID("1234567890").
		IIN("123456789012").
		PhoneNumber("1234567890").
		CreatedAt(time.Now()).
		V()

	querySelect := regexp.QuoteMeta(
		`SELECT
			id, first_name, last_name, middle_name, address, phone_number, email, role, password, birth_date, iin, government_id, created_at
		FROM
			public.users	
		WHERE
			id = $1`,
	)

	args := []interface{}{
		v.ID,
	}

	rows := pgxmock.NewRows([]string{
		"id",
		"first_name",
		"last_name",
		"middle_name",
		"address",
		"phone_number",
		"email",
		"role",
		"password",
		"birth_date",
		"iin",
		"government_id",
		"created_at",
	}).
		AddRow(
			v.ID,
			v.FirstName,
			v.LastName,
			v.MiddleName,
			v.Address,
			v.PhoneNumber,
			v.Email,
			v.Role,
			v.Password,
			v.BirthDate,
			v.IIN,
			v.GovernmentID,
			v.CreatedAt,
		)

	f.pgxPoolMock.ExpectQuery(querySelect).
		WithArgs(args...).
		WillReturnRows(rows)

	ctx := context.Background()

	user, err := f.storage.UserGetByID(ctx, v.ID)

	expectErr := f.pgxPoolMock.ExpectationsWereMet()

	// assert
	require.NoError(t, err)
	require.NoError(t, expectErr)

	require.NotEmpty(t, user)

	assert.Equal(t, v.ID, user.ID)
	assert.Equal(t, v.FirstName, user.FirstName)
	assert.Equal(t, v.LastName, user.LastName)
	assert.Equal(t, v.MiddleName, user.MiddleName)
	assert.Equal(t, v.PhoneNumber, user.PhoneNumber)
	assert.Equal(t, v.IIN, user.IIN)
	assert.Equal(t, v.GovernmentID, user.GovernmentID)
	assert.Equal(t, v.Address, user.Address)
	assert.Equal(t, v.Email, user.Email)
	assert.Equal(t, v.Role, user.Role)
	assert.Equal(t, v.BirthDate, user.BirthDate)
	assert.Equal(t, v.CreatedAt, user.CreatedAt)
}
