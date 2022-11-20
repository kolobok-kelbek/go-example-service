package infra

import (
	"github.com/kolobok-kelbek/cong"
	"github.com/kolobok-kelbek/go-example-service/infra/config"
	"github.com/kolobok-kelbek/go-example-service/static"
	"net/http"
)

type App struct {
	config *config.Config
}

func NewApp() *App {
	// TODO: logger
	app := &App{}

	loader := cong.NewLoader[config.Config]()
	cfg, err := loader.LoadFromEmbedFSByPath("ExampleService", static.Snapshot, "config", cong.YamlExt)
	if err != nil {
		panic(err)
	}

	app.config = cfg

	return app
}

func (app *App) Run() {
	// TODO: routing
	http.HandleFunc("/hello", app.getHello)

	err := http.ListenAndServe(app.config.Server.Port, nil)
	if err != nil {
		panic(err)
	}
}

func (app *App) getHello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		return
	}
}
