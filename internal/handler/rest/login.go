package rest

import "net/http"

type LoginHandler struct {
}

func NewLoginHandler() (*LoginHandler, error) {
	return &LoginHandler{}, nil
}

func (h *LoginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}
