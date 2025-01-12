package test

import (
	configDir "github.com/kolobok-kelbek/tomato/config"
	"github.com/kolobok-kelbek/tomato/internal/app"
	"github.com/kolobok-kelbek/tomato/internal/config"
	"net/http"
)

func InitRoute() (*http.ServeMux, error) {
	cfg, err := config.Load(configDir.Snapshot, config.TestEnv)
	if err != nil {
		panic(err)
	}

	di, err := app.NewDIContainer(cfg)
	if err != nil {
		return nil, err
	}

	return di.GetRoute(), nil
}
