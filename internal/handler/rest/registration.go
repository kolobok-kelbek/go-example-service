package rest

import "net/http"

type RegistrationHandler struct {
}

func NewRegistrationHandler() (*RegistrationHandler, error) {
	return &RegistrationHandler{}, nil
}

func (h *RegistrationHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Registration"))
}
