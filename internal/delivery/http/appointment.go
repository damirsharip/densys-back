package http

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/densys/admin/internal/entity"
)

type appointmentCreateRequest struct {
	ServiceId string `json:"service_id,omitempty"`
	DoctorId  string `json:"doctor_id"`
	Date      string `json:"date"`
	Time      string `json:"time"`
}

type appointmentCreateResponse struct {
	Appointment entity.Appointment `json:"appointment"`
}

// appointmentCreate creates a new appointment in the system handler.
func (h *handler) appointmentCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req appointmentCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		date, errD := time.Parse("2006-01-02", req.Date)
		if errD != nil {
			respondError(w, http.StatusBadRequest, "invalid date format")
			return
		}

		scheduleTime, errT := time.Parse("15:04", req.Time)
		if errT != nil {
			respondError(w, http.StatusBadRequest, "invalid time format")
			return
		}

		payload, err := getUserPayload(r)
		if err != nil {
			respondError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		appointment, err := h.service.AppointmentCreate(r.Context(), entity.Appointment{
			Patient: entity.User{ID: payload.UserID},
			Doctor:  entity.User{ID: req.DoctorId},
			Service: entity.Service{ID: req.ServiceId},
			Date:    date,
			Time:    scheduleTime,
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println(err)
			return
		}

		respondJSON(w, http.StatusCreated, appointmentCreateResponse{Appointment: appointment})
	})
}

type appointmentGetAllResponse struct {
	Appointment []entity.Appointment `json:"appointments"`
}

// appointmentServiceGetAll gets all appointments from the system handler.
func (h *handler) appointmentServiceGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appointments, err := h.service.AppointmentServiceGetAll(r.Context())

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, appointmentGetAllResponse{Appointment: appointments})
	})
}

// appointmentDoctorGetAll gets all appointments from the system handler.
func (h *handler) appointmentDoctorGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appointments, err := h.service.AppointmentDoctorGetAll(r.Context())

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println(err)
			return
		}

		respondJSON(w, http.StatusOK, appointmentGetAllResponse{Appointment: appointments})
	})
}

// appointmentOfPatientGetAll gets all appointments of patient from the system handler.
func (h *handler) appointmentServiceOfPatientGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := getUserPayload(r)
		if err != nil {
			respondError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		appointments, err := h.service.AppointmentServiceGetAllOfPatient(r.Context(), payload.UserID)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusCreated, appointmentGetAllResponse{Appointment: appointments})
	})
}

// appointmentOfPatientGetAll gets all appointments of patient from the system handler.
func (h *handler) appointmentDoctorOfPatientGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := getUserPayload(r)
		if err != nil {
			respondError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		appointments, err := h.service.AppointmentDoctorGetAllOfPatient(r.Context(), payload.UserID)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusCreated, appointmentGetAllResponse{Appointment: appointments})
	})
}

// appointmentOfDoctorGetAll gets all appointments of doctor from the system handler.
func (h *handler) appointmentOfDoctorGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		doctorId := httprouter.ParamsFromContext(r.Context()).ByName("doctorId")

		if _, err := uuid.Parse(doctorId); err != nil {
			respondError(w, http.StatusBadRequest, "invalid doctor id")
			return
		}

		appointments, err := h.service.AppointmentGetAllOfDoctor(r.Context(), doctorId)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusCreated, appointmentGetAllResponse{Appointment: appointments})
	})
}

// appointmentOfPatientGetAll gets all appointments of doctor from the system handler.
func (h *handler) appointmentOfPatientGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := getUserPayload(r)
		if err != nil {
			respondError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		appointments, err := h.service.AppointmentGetAllOfPatient(r.Context(), payload.UserID)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusCreated, appointmentGetAllResponse{Appointment: appointments})
	})
}

// appointmentOfDoctorGetAll gets all appointments of doctor from the system handler.
func (h *handler) appointmentOfServiceGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serviceId := httprouter.ParamsFromContext(r.Context()).ByName("serviceId")

		if _, err := uuid.Parse(serviceId); err != nil {
			respondError(w, http.StatusBadRequest, "invalid service id")
			return
		}

		appointments, err := h.service.AppointmentGetAllOfService(r.Context(), serviceId)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusCreated, appointmentGetAllResponse{Appointment: appointments})
	})
}

// appointmentGetAll gets all doctors from the system handler.
func (h *handler) appointmentApprove() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := httprouter.ParamsFromContext(r.Context()).ByName("id")

		if _, err := uuid.Parse(id); err != nil {
			respondError(w, http.StatusBadRequest, "invalid patient id")
			return
		}

		appointment, err := h.service.AppointmentApprove(r.Context(), id)

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusOK, appointmentCreateResponse{Appointment: appointment})
	})
}
