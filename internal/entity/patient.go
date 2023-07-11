package entity

const (
	RolePatient = "patient"
)

// Patient represents a patient.
type Patient struct {
	ID                     string `json:"id"`
	User                   User   `json:"user"`
	BloodGroup             string `json:"blood_group"`
	EmergencyContactNumber string `json:"emergency_contact_number"`
	MaritalStatus          string `json:"marital_status"`
}

// PatientUpdateInput represents DTO for updating patient.
type PatientUpdateInput struct {
	BloodGroup             *string `json:"blood_group,omitempty"`
	EmergencyContactNumber *string `json:"emergency_contact_number,omitempty"`
	MaritalStatus          *string `json:"marital_status,omitempty"`
}
