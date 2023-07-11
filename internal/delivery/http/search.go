package http

import (
	"log"
	"net/http"
)

type autoCompleteResponse struct {
	Result []string `json:"result"`
}

// autoComplete gets auto complete result.
func (h *handler) autoComplete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		autocomplete, err := h.service.AutoComplete(r.Context())

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println(err)
			return
		}

		respondJSON(w, http.StatusOK, autoCompleteResponse{Result: autocomplete})
	})
}
