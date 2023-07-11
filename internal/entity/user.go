package entity

import "time"

// User represents a user info in the system. It is general for all users.
type User struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"firstName,omitempty"`
	LastName     string    `json:"lastName,omitempty"`
	MiddleName   string    `json:"middleName,omitempty"`
	Address      string    `json:"address,omitempty"`
	Email        string    `json:"email,omitempty"`
	Role         string    `json:"role,omitempty"`
	Password     string    `json:"-"`
	BirthDate    time.Time `json:"birthDate,omitempty"`
	PhoneNumber  string    `json:"phoneNumber,omitempty"`
	IIN          string    `json:"iin,omitempty"`
	GovernmentID string    `json:"government_id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

// UserUpdateInput represents DTO for updating user.
type UserUpdateInput struct {
	FirstName    *string    `json:"firstName,omitempty"`
	LastName     *string    `json:"lastName,omitempty"`
	MiddleName   *string    `json:"middleName,omitempty"`
	Address      *string    `json:"address,omitempty"`
	Email        *string    `json:"email,omitempty"`
	BirthDate    *time.Time `json:"birthDate,omitempty"`
	PhoneNumber  *string    `json:"phoneNumber,omitempty"`
	IIN          *string    `json:"iin,omitempty"`
	GovernmentID *string    `json:"government_id,omitempty"`
}
