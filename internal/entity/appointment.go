package entity

import "time"

// Appointment represents an appointment information.
type Appointment struct {
	ID         string    `json:"id"`
	Patient    User      `json:"patient_id,omitempty"`
	Doctor     User      `json:"doctor_id,omitempty"`
	Service    Service   `json:"service_id,omitempty"`
	Date       time.Time `json:"date"`
	Time       time.Time `json:"time"`
	IsApproved bool      `json:"is_approved"`
}
