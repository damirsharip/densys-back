package http

import (
	"net/http"

	"github.com/khanfromasia/densys/admin/internal/entity"
)

type specializationCreateRequest struct {
	Name string `json:"name"`
}

type specializationCreateResponse struct {
	Specialization entity.Specialization `json:"specialization"`
}

// specializationCreate creates a new specialization in the system handler.
func (h *handler) specializationCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req specializationCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		specialization, err := h.service.SpecializationCreate(r.Context(), entity.Specialization{Name: req.Name})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, specializationCreateResponse{Specialization: specialization})
	})
}

type specializationGetAllResponse struct {
	Specializations []entity.Specialization `json:"specializations"`
}

// specializationGetAll gets all specializations in the system handler.
func (h *handler) specializationGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		specializations, err := h.service.SpecializationGetAll(r.Context(), name)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, specializationGetAllResponse{Specializations: specializations})
	})
}
