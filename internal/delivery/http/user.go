package http

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/densys/admin/internal/entity"
)

type userUpdateRequest struct {
	FirstName    *string    `json:"firstName,omitempty"`
	LastName     *string    `json:"lastName,omitempty"`
	MiddleName   *string    `json:"middleName,omitempty"`
	Address      *string    `json:"address,omitempty"`
	Email        *string    `json:"email,omitempty"`
	BirthDate    *time.Time `json:"birthDate,omitempty"`
	PhoneNumber  *string    `json:"phoneNumber,omitempty"`
	IIN          *string    `json:"iin,omitempty"`
	GovernmentID *string    `json:"government_id,omitempty"`
}

type userUpdateResponse struct {
	User entity.User `json:"user"`
}

// userUpdate updates user info in the system handler.
func (h *handler) userUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := httprouter.ParamsFromContext(r.Context()).ByName("userId")

		if _, err := uuid.Parse(userId); err != nil {
			respondError(w, http.StatusBadRequest, "invalid user id")
			return
		}

		var req userUpdateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		user, err := h.service.UserUpdate(r.Context(), userId, entity.UserUpdateInput{
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			MiddleName:   req.MiddleName,
			Address:      req.Address,
			Email:        req.Email,
			BirthDate:    req.BirthDate,
			PhoneNumber:  req.PhoneNumber,
			IIN:          req.IIN,
			GovernmentID: req.GovernmentID,
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "failed to update user")
			return
		}

		respondJSON(w, http.StatusOK, userUpdateResponse{User: user})
	})
}
