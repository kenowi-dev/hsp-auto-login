package api

import (
	"net/http"
	"svelte-go/hspscraper"
	"time"
)

type RegisterRequest struct {
	Course   hspscraper.Course `json:"course"`
	Sport    string            `json:"sport"`
	Date     time.Time         `json:"date"`
	Email    string            `json:"email"`
	Password string            `json:"password"`
}

func RegisterHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest
	err := parseBody(r, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = hspscraper.Register(request.Course, request.Sport, request.Email, request.Password, request.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
