package dialect

import (
	"github.com/kolobok-kelbek/tomato/internal/db/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PgProduce(config config.DataBase) gorm.Dialector {
	dsn := ""
	if config.DSN == "" {
		dsn = config.Dsn()
	} else {
		dsn = config.DSN
	}

	return postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: config.PreferSimpleProtocol,
	})
}
