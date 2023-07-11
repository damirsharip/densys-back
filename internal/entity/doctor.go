package entity

const (
	RoleDoctor = "doctor"
)

// Doctor represents a doctor info in the system.
type Doctor struct {
	ID               string         `json:"id"`
	User             User           `json:"user,omitempty"`
	Department       Department     `json:"department,omitempty"`
	Specialization   Specialization `json:"specialization,omitempty"`
	WorkExperience   int            `json:"workExperience,omitempty"`
	Photo            string         `json:"photo,omitempty"`
	Category         string         `json:"category,omitempty"`
	ScheduleDetails  string         `json:"scheduleDetails,omitempty"`
	Degree           string         `json:"degree,omitempty"`
	AppointmentPrice int            `json:"appointmentPrice,omitempty"`
	Rating           int            `json:"rating,omitempty"`
	Services         []Service      `json:"services"`
}

// DoctorUpdateInput represents DTO for updating doctor.
type DoctorUpdateInput struct {
	DepartmentID     *string `json:"departmentId,omitempty"`
	SpecializationID *string `json:"specializationId,omitempty"`
	WorkExperience   *int    `json:"workExperience,omitempty"`
	Photo            *string `json:"photo,omitempty"`
	Category         *string `json:"category,omitempty"`
	ScheduleDetails  *string `json:"scheduleDetails,omitempty"`
	Degree           *string `json:"degree,omitempty"`
	AppointmentPrice *int    `json:"appointmentPrice,omitempty"`
	Rating           *int    `json:"rating,omitempty"`
}

// DoctorWithSchedule represents a doctor info in the system.
type DoctorWithSchedule struct {
	ID               string         `json:"id"`
	User             User           `json:"user"`
	Department       Department     `json:"department"`
	Specialization   Specialization `json:"specialization"`
	WorkExperience   int            `json:"workExperience"`
	Photo            string         `json:"photo"`
	Category         string         `json:"category"`
	ScheduleDetails  string         `json:"scheduleDetails"`
	Degree           string         `json:"degree"`
	AppointmentPrice int            `json:"appointmentPrice"`
	Rating           int            `json:"rating"`
	Schedule         []string       `json:"schedule"`
}

type Schedule string
