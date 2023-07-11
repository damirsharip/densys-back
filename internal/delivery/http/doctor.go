package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/densys/admin/internal/entity"
)

type doctorCreateRequest struct {
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
	Doctor struct {
		DepartmentID     string `json:"departmentId"`
		SpecializationID string `json:"specializationId"`
		WorkExperience   int    `json:"workExperience"`
		Photo            string `json:"photo"`
		Category         string `json:"category"`
		ScheduleDetails  string `json:"scheduleDetails"`
		Degree           string `json:"degree"`
		AppointmentPrice int    `json:"appointmentPrice"`
		Rating           int    `json:"rating"`
	} `json:"doctor"`
}

type doctorCreateResponse struct {
	Doctor entity.Doctor `json:"doctor"`
}

// doctorCreate creates a new doctor in the system handler.
func (h *handler) doctorCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req doctorCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		doctor, err := h.service.DoctorCreate(r.Context(), entity.Doctor{
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
			Department: entity.Department{
				ID: req.Doctor.DepartmentID,
			},
			Specialization: entity.Specialization{
				ID: req.Doctor.SpecializationID,
			},
			WorkExperience:   req.Doctor.WorkExperience,
			Photo:            req.Doctor.Photo,
			Category:         req.Doctor.Category,
			ScheduleDetails:  req.Doctor.ScheduleDetails,
			Degree:           req.Doctor.Degree,
			AppointmentPrice: req.Doctor.AppointmentPrice,
			Rating:           req.Doctor.Rating,
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusCreated, doctorCreateResponse{Doctor: doctor})
	})
}

type doctorUpdateRequest struct {
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

type doctorUpdateResponse struct {
	Doctor entity.Doctor `json:"doctor"`
}

// doctorUpdate updates doctor info in the system handler.
func (h *handler) doctorUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := httprouter.ParamsFromContext(r.Context()).ByName("userId")

		if _, err := uuid.Parse(userId); err != nil {
			respondError(w, http.StatusBadRequest, "invalid user id")
			return
		}

		var req doctorUpdateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		doctor, err := h.service.DoctorUpdate(r.Context(), userId, entity.DoctorUpdateInput{
			DepartmentID:     req.DepartmentID,
			SpecializationID: req.SpecializationID,
			WorkExperience:   req.WorkExperience,
			Photo:            req.Photo,
			Category:         req.Category,
			ScheduleDetails:  req.ScheduleDetails,
			Degree:           req.Degree,
			AppointmentPrice: req.AppointmentPrice,
			Rating:           req.Rating,
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, doctorUpdateResponse{Doctor: doctor})
	})
}

type doctorGetAllResponse struct {
	Doctors []entity.Doctor `json:"doctors"`
}

// doctorGetAll gets all doctors from the system handler.
func (h *handler) doctorGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		doctors, err := h.service.DoctorGetAll(r.Context(), name)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, doctorGetAllResponse{Doctors: doctors})
	})
}

type doctorSearchByServiceAndSpecializationResponse struct {
	Doctors    []entity.Doctor   `json:"doctors"`
	Pagination entity.Pagination `json:"pagination"`
}

// doctorSearchByServiceAndSpecialization gets all doctors from the system handler.
func (h *handler) doctorSearchByServiceAndSpecialization() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		specializationId := r.URL.Query().Get("specializationId")
		pageNumber := r.URL.Query().Get("pageNumber")

		// parse page number
		page, err := strconv.ParseInt(pageNumber, 10, 64)

		if err != nil {
			page = 1
		}

		doctors, pagination, err := h.service.DoctorSearchByServiceAndSpecialization(r.Context(), query, specializationId, entity.Pagination{CurrentPage: page})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, doctorSearchByServiceAndSpecializationResponse{
			Doctors:    doctors,
			Pagination: pagination,
		})
	})
}

// doctorGetByUserID gets doctor by user id from the system handler.
func (h *handler) doctorGetByUserID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := httprouter.ParamsFromContext(r.Context()).ByName("userId")

		if _, err := uuid.Parse(userId); err != nil {
			respondError(w, http.StatusBadRequest, "invalid user id")
			return
		}

		doctor, err := h.service.DoctorGetByUserID(r.Context(), userId)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, doctor)
	})
}

type doctorGetAllBySpecializationResponse struct {
	Doctors     []entity.DoctorWithSchedule `json:"doctors"`
	CurrentPage int                         `json:"current_page"`
}

// doctorGetAllBySpecialization gets all doctors with particular specialization from the system handler.
func (h *handler) doctorGetAllBySpecialization() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := httprouter.ParamsFromContext(r.Context()).ByName("id")
		date := httprouter.ParamsFromContext(r.Context()).ByName("date")

		if _, err := uuid.Parse(id); err != nil {
			respondError(w, http.StatusBadRequest, "invalid specialization id")
			return
		}

		//date.(time.Time).Day()
		page := r.URL.Query().Get("page_number")
		pageNumber, err := strconv.Atoi(page)
		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		doctors, err := h.service.DoctorGetAllBySpecialization(r.Context(), id, date, pageNumber)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, doctorGetAllBySpecializationResponse{Doctors: doctors, CurrentPage: pageNumber})
	})
}

type doctorGetScheduleByDate struct {
	Schedule []entity.Schedule `json:"schedule"`
}

// doctorGetAllBySpecialization gets all doctors with particular specialization from the system handler.
func (h *handler) doctorGetScheduleByDate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := httprouter.ParamsFromContext(r.Context()).ByName("userId")
		date := httprouter.ParamsFromContext(r.Context()).ByName("date")

		if _, err := uuid.Parse(userId); err != nil {
			respondError(w, http.StatusBadRequest, "invalid specialization id")
			return
		}

		schedule, err := h.service.DoctorGetScheduleByDate(r.Context(), userId, date)

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		respondJSON(w, http.StatusOK, doctorGetScheduleByDate{Schedule: schedule})
	})
}
