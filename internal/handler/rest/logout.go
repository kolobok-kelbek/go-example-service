package rest

import "net/http"

type LogoutHandler struct {
}

func NewLogoutHandler() (*LogoutHandler, error) {
	return &LogoutHandler{}, nil
}

func (h *LogoutHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
