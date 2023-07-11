package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/densys/admin/internal/service"
)

const (
	v0 = "/api/v0/admin"
	v1 = "/api/v0"
)

type handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) SetupRoutes() *httprouter.Router {
	router := httprouter.New()

	router.Handler(http.MethodGet, v0+"/specializations", h.authMiddleware(h.specializationGetAll()))
	router.Handler(http.MethodPost, v0+"/specializations", h.authMiddleware(h.specializationCreate()))
	router.Handler(http.MethodGet, v1+"/specializations/:id/doctors/:date", h.authMiddleware(h.doctorGetAllBySpecialization()))

	router.Handler(http.MethodGet, v0+"/departments", h.authMiddleware(h.departmentGetAll()))
	router.Handler(http.MethodPost, v0+"/departments", h.authMiddleware(h.departmentCreate()))

	router.Handler(http.MethodGet, v1+"/schedule/:userId/:date", h.authMiddleware(h.doctorGetScheduleByDate()))
	router.Handler(http.MethodGet, v0+"/services", h.authMiddleware(h.serviceGetAll()))
	router.Handler(http.MethodPost, v0+"/services", h.authMiddleware(h.serviceCreate()))

	router.Handler(http.MethodGet, v1+"/search", h.authMiddleware(h.doctorSearchByServiceAndSpecialization()))
	router.Handler(http.MethodGet, v1+"/autocomplete", h.authMiddleware(h.autoComplete()))

	router.Handler(http.MethodGet, v0+"/doctors", h.authMiddleware(h.doctorGetAll()))
	router.Handler(http.MethodPost, v0+"/doctors", h.authMiddleware(h.doctorCreate()))
	router.Handler(http.MethodPut, v0+"/doctors/:userId", h.authMiddleware(h.doctorUpdate()))
	router.Handler(http.MethodGet, v0+"/doctors/:userId", h.authMiddleware(h.doctorGetByUserID()))

	router.Handler(http.MethodGet, v1+"/doctors/:userId/services", h.authMiddleware(h.doctorServiceGetAll()))

	router.Handler(http.MethodGet, v0+"/patients", h.authMiddleware(h.patientGetAll()))
	router.Handler(http.MethodPost, v0+"/patients", h.authMiddleware(h.patientCreate()))
	router.Handler(http.MethodPut, v0+"/patients/:userId", h.authMiddleware(h.patientUpdate()))
	router.Handler(http.MethodGet, v0+"/patients/:userId", h.authMiddleware(h.patientGetByUserID()))

	router.Handler(http.MethodGet, v0+"/appointments/service", h.authMiddleware(h.appointmentServiceGetAll()))
	router.Handler(http.MethodGet, v0+"/appointments/doctor", h.authMiddleware(h.appointmentDoctorGetAll()))
	router.Handler(http.MethodPut, v0+"/appointments/:id/approve", h.authMiddleware(h.appointmentApprove()))

	router.Handler(http.MethodPost, v1+"/appointments", h.authMiddleware(h.appointmentCreate()))
	router.Handler(http.MethodGet, v1+"/appointments/service", h.authMiddleware(h.appointmentServiceOfPatientGetAll()))
	router.Handler(http.MethodGet, v1+"/appointments/doctor", h.authMiddleware(h.appointmentDoctorOfPatientGetAll()))
	router.Handler(http.MethodGet, v1+"/appointments/service/:serviceId", h.authMiddleware(h.appointmentOfServiceGetAll()))
	router.Handler(http.MethodGet, v1+"/appointments/doctor/:doctorId", h.authMiddleware(h.appointmentOfDoctorGetAll()))
	router.Handler(http.MethodGet, v1+"/appointments", h.authMiddleware(h.appointmentOfPatientGetAll()))

	router.Handler(http.MethodPut, v0+"/users/:userId", h.authMiddleware(h.userUpdate()))

	router.Handler(http.MethodPost, v1+"/signin", h.signIn())

	return router
}
