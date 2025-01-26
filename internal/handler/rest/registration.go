package rest

import (
	"encoding/json"
	"github.com/kolobok-kelbek/tomato/internal/domain/auth"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type RegistrationRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegistrationHandler struct {
	db *gorm.DB
}

func NewRegistrationHandler(db *gorm.DB) (*RegistrationHandler, error) {
	return &RegistrationHandler{
		db: db,
	}, nil
}

func (h *RegistrationHandler) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	var req RegistrationRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	err = h.db.Create(&auth.Login{Login: req.Login, Password: req.Password}).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
