package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/densys/admin/internal/entity"
)

type serviceCreateRequest struct {
	Name                    string `json:"name"`
	Price                   string `json:"price"`
	ListOfContradictions    string `json:"list_of_contradictions"`
	OtherRelatedInformation string `json:"other_related_information"`
	SpecializationID        string `json:"specialization_id"`
}

type serviceCreateResponse struct {
	Service entity.Service `json:"service"`
}

// departmentCreate performs the create department operation handler.
func (h *handler) serviceCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req serviceCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		arg := entity.Service{
			Name:                    req.Name,
			Price:                   req.Price,
			ListOfContradictions:    req.ListOfContradictions,
			OtherRelatedInformation: req.OtherRelatedInformation,
			Specialization: entity.Specialization{
				ID: req.SpecializationID,
			},
		}

		department, err := h.service.ServiceCreate(r.Context(), arg)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, serviceCreateResponse{Service: department})
	})
}

type serviceGetAllResponse struct {
	Services []entity.Service `json:"services"`
}

// departmentGetAll performs the get all departments operation handler.
func (h *handler) serviceGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		departments, err := h.service.ServiceGetAll(r.Context())

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, serviceGetAllResponse{Services: departments})
	})
}

func (h *handler) doctorServiceGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := httprouter.ParamsFromContext(r.Context()).ByName("userId")
		services, err := h.service.DoctorServiceGetAll(r.Context(), userId)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, serviceGetAllResponse{Services: services})
	})
}
