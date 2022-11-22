package database

import (
	"context"
	"github.com/kolobok-kelbek/tomato/pkg/db/config"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB     *gorm.DB
	Config config.DataBase
	Model  Model
}

func (repository *GormRepository) GetDB() *gorm.DB {
	return repository.DB
}

func (repository *GormRepository) Begin() (*gorm.DB, context.CancelFunc) {
	ctx, cancel := repository.GetContextTimeout()

	return repository.BeginCtx(ctx), cancel
}

func (repository *GormRepository) BeginCtx(ctx context.Context) *gorm.DB {
	return repository.DB.Model(repository.GetModel()).WithContext(ctx)
}

func (repository *GormRepository) GetModel() Model {
	return repository.Model
}

func (repository *GormRepository) GetContextTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), repository.Config.GetTimeout())
}
