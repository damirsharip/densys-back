package http

import (
	"log"
	"net/http"

	"github.com/khanfromasia/densys/admin/internal/entity"
)

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	Token entity.Token `json:"token"`
	User  entity.User  `json:"user"`
}

// signIn performs the sign in operation http.Handler.
func (h *handler) signIn() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req signInRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		user, token, err := h.service.SignIn(r.Context(), req.Email, req.Password)
		if err != nil {
			respondError(w, http.StatusUnauthorized, "invalid credentials")
			log.Println(err)
			return
		}

		respondJSON(w, http.StatusOK, signInResponse{
			Token: token,
			User:  user,
		})
	})
}
