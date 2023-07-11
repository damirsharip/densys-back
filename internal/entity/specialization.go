package entity

// Specialization represents a hospital specialization.
type Specialization struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}
