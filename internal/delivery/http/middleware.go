package http

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/khanfromasia/densys/admin/internal/pkg/jwt"
)

// `authMiddleware is a middleware that checks if the user is authenticated`
func (h *handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request header
		token := r.Header.Get("Authorization")
		if token == "" {
			respondError(w, http.StatusUnauthorized, "missing token")
			return
		}

		t := strings.Split(token, " ")

		// Verify the token
		payload, err := h.service.TokenVerify(r.Context(), t[1])
		if err != nil {
			respondError(w, http.StatusUnauthorized, err.Error())
			return
		}

		// Add the payload to the request context
		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserPayload(r *http.Request) (jwt.Payload, error) {
	id := r.Context().Value("payload")

	payload, ok := id.(jwt.Payload)
	if !ok {
		return jwt.Payload{}, errors.New("user payload is of invalid type")
	}

	return payload, nil
}
