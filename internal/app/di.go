package app

import (
	"github.com/kolobok-kelbek/tomato/internal/config"
	database "github.com/kolobok-kelbek/tomato/internal/db"
	"github.com/kolobok-kelbek/tomato/internal/handler/rest"
	"github.com/kolobok-kelbek/tomato/internal/logger"
	"net/http"
)

type DIContainer struct {
	app   *App
	route *http.ServeMux
}

func NewDIContainer(cfg *config.Config) (*DIContainer, error) {
	source, err := database.GetDataSource(cfg.DataBase)
	if err != nil {
		return nil, err
	}

	login, err := rest.NewLoginHandler()
	if err != nil {
		return nil, err
	}
	logout, err := rest.NewLogoutHandler()
	if err != nil {
		return nil, err
	}
	registration, err := rest.NewRegistrationHandler(source)
	if err != nil {
		return nil, err
	}

	lg := logger.NewLogger()
	route := rest.NewRoute(login, logout, registration)
	server := NewServer(cfg.Server, route)
	app := NewApp(cfg, server, lg)
	return &DIContainer{
		app:   app,
		route: route,
	}, nil
}

func (c *DIContainer) App() *App {
	return c.app
}

func (c *DIContainer) Route() *http.ServeMux {
	return c.route
}
