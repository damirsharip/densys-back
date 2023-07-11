package entity

// Service represents service information.
type Service struct {
	ID                      string         `json:"id"`
	Name                    string         `json:"name"`
	Price                   string         `json:"price"`
	ListOfContradictions    string         `json:"list_of_contradictions"`
	OtherRelatedInformation string         `json:"other_related_information"`
	Specialization          Specialization `json:"specialization,omitempty"`
}
