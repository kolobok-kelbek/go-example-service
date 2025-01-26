package dialect

import (
	"github.com/go-task/task/v3/errors"
	"github.com/kolobok-kelbek/tomato/internal/db/config"
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

func (p *Manager) Produce(config config.DataBase) (gorm.Dialector, error) {
	if producer, ok := p.producers[config.Dialect]; ok {
		return producer(config), nil
	} else {
		return nil, errors.New("incorrect dialect: " + config.Dialect)
	}
}
