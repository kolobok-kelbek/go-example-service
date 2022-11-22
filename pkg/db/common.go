package database

import (
	"context"

	"gorm.io/gorm"
)

type Model interface {
	TableName() string
	CreateNew() Model
}

type Repository interface {
	GetDB() *gorm.DB
	Begin() (*gorm.DB, context.CancelFunc)
	BeginCtx(ctx context.Context) *gorm.DB
	GetModel() Model
	GetContextTimeout() (context.Context, context.CancelFunc)
}
