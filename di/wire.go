//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kolobok-kelbek/go-example-service/infra"
)

func Init() *infra.App {
	wire.Build(infra.NewApp)

	return &infra.App{}
}
