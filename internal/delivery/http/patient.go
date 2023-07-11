package http

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/densys/admin/internal/entity"
)

type patientCreateRequest struct {
	User struct {
		FirstName    string    `json:"firstName"`
		LastName     string    `json:"lastName"`
		MiddleName   string    `json:"middleName"`
		Email        string    `json:"email"`
		Address      string    `json:"address"`
		PhoneNumber  string    `json:"phoneNumber"`
		IIN          string    `json:"iin"`
		GovernmentID string    `json:"governmentId"`
		Password     string    `json:"password"`
		BirthDate    time.Time `json:"birthDate"`
	} `json:"user"`
	Patient struct {
		BloodGroup             string `json:"bloodGroup"`
		EmergencyContactNumber string `json:"emergencyContactNumber"`
		MaritalStatus          string `json:"maritalStatus"`
	} `json:"patient"`
}

type patientCreateResponse struct {
	Patient entity.Patient `json:"patient"`
}

// patientCreate creates a new patient in the system handler.
func (h *handler) patientCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req patientCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		patient, err := h.service.PatientCreate(r.Context(), entity.Patient{
			User: entity.User{
				FirstName:    req.User.FirstName,
				LastName:     req.User.LastName,
				MiddleName:   req.User.MiddleName,
				Email:        req.User.Email,
				Address:      req.User.Address,
				PhoneNumber:  req.User.PhoneNumber,
				IIN:          req.User.IIN,
				GovernmentID: req.User.GovernmentID,
				BirthDate:    req.User.BirthDate,
				Password:     req.User.Password,
			},
			BloodGroup:             req.Patient.BloodGroup,
			EmergencyContactNumber: req.Patient.EmergencyContactNumber,
			MaritalStatus:          req.Patient.MaritalStatus,
		})
		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusCreated, patientCreateResponse{Patient: patient})
	})
}

type patientUpdateRequest struct {
	BloodGroup             *string `json:"blood_group,omitempty"`
	EmergencyContactNumber *string `json:"emergency_contact_number,omitempty"`
	MaritalStatus          *string `json:"marital_status,omitempty"`
}

type patientUpdateResponse struct {
	Patient entity.Patient `json:"patient"`
}

// patientUpdate updates a patient in the system handler.
func (h *handler) patientUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := httprouter.ParamsFromContext(r.Context()).ByName("userId")

		if _, err := uuid.Parse(userId); err != nil {
			respondError(w, http.StatusBadRequest, "invalid user id")
			return
		}

		var req patientUpdateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		patient, err := h.service.PatientUpdate(r.Context(), userId, entity.PatientUpdateInput{
			BloodGroup:             req.BloodGroup,
			EmergencyContactNumber: req.EmergencyContactNumber,
			MaritalStatus:          req.MaritalStatus,
		})
		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, patientUpdateResponse{Patient: patient})
	})
}

type patientGetAllResponse struct {
	Patients []entity.Patient `json:"patients"`
}

// patientGetAll gets all patients in the system handler.
func (h *handler) patientGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		patients, err := h.service.PatientGetAll(r.Context())
		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, patientGetAllResponse{Patients: patients})
	})
}

// patientGetByUserID gets a patient by user id in the system handler.
func (h *handler) patientGetByUserID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := httprouter.ParamsFromContext(r.Context()).ByName("userId")

		if _, err := uuid.Parse(userId); err != nil {
			respondError(w, http.StatusBadRequest, "invalid user id")
			return
		}

		patient, err := h.service.PatientGetByUserID(r.Context(), userId)
		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, patient)
	})
}
