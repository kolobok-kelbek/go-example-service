package database

import (
	"github.com/kolobok-kelbek/tomato/internal/db/config"
	"github.com/kolobok-kelbek/tomato/internal/db/dialect"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func GetDataSource(cfg config.DataBase, dialectManager *dialect.Manager) *gorm.DB {
	if db != nil {
		return db
	}

	dialector := dialectManager.Produce(cfg)

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second * time.Duration(cfg.SlowThreshold),
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)

	newDb, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: false,
		Logger:      gormLogger,
	})

	if err != nil {
		panic(err)
	}

	db = newDb

	return newDb
}
