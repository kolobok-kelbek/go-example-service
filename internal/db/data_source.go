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

func GetDataSource(cfg config.DataBase) (*gorm.DB, error) {
	dialectManager := dialect.NewManager()
	dialector, err := dialectManager.Produce(cfg)
	if err != nil {
		return nil, err
	}

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second * time.Duration(cfg.SlowThreshold),
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: false,
		Logger:      gormLogger,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
