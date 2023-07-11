package http

import (
	"net/http"

	"github.com/khanfromasia/densys/admin/internal/entity"
)

type departmentCreateRequest struct {
	Name string `json:"name"`
}

type departmentCreateResponse struct {
	Department entity.Department `json:"department"`
}

// departmentCreate performs the create department operation handler.
func (h *handler) departmentCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req departmentCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		department, err := h.service.DepartmentCreate(r.Context(), entity.Department{Name: req.Name})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, departmentCreateResponse{Department: department})
	})
}

type departmentGetAllResponse struct {
	Departments []entity.Department `json:"departments"`
}

// departmentGetAll performs the get all departments operation handler.
func (h *handler) departmentGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		departments, err := h.service.DepartmentGetAll(r.Context())

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, departmentGetAllResponse{Departments: departments})
	})
}
