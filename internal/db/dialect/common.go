package dialect

import (
	"errors"
	"github.com/kolobok-kelbek/tomato/internal/db/config"
	"log"

	"gorm.io/gorm"
)

const postgresDialectName = "postgres"

type Produce func(config config.DataBase) gorm.Dialector

type Manager struct {
	producers map[string]Produce
}

func NewManager() *Manager {
	producer := Manager{producers: map[string]Produce{
		postgresDialectName: PgProduce,
	}}

	return &producer
}

func (p *Manager) Produce(config config.DataBase) gorm.Dialector {
	if producer, ok := p.producers[config.Dialect]; ok {
		return producer(config)
	} else {
		message := "incorrect dialect: " + config.Dialect
		log.Fatal(message)
		panic(errors.New(message)) //nolint:goerr113
	}
}
