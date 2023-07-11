package entity

type Pagination struct {
	TotalCount  int64 `json:"totalCount"`
	TotalPages  int64 `json:"totalPages"`
	CurrentPage int64 `json:"currentPage"`
	PerPage     int64 `json:"perPage"`
}
