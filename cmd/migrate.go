package cmd

import (
	configFiles "github.com/kolobok-kelbek/tomato/config"
	"github.com/kolobok-kelbek/tomato/internal/config"
	"github.com/kolobok-kelbek/tomato/internal/db"
	"github.com/kolobok-kelbek/tomato/migrations"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		goose.SetBaseFS(migrations.Snapshot)

		cfg, err := config.Load(configFiles.Snapshot, config.ProdEnv)
		if err != nil {
			panic(err)
		}

		err = goose.SetDialect(cfg.DataBase.Dialect)
		if err != nil {
			panic(err)
		}

		gormDB, err := database.GetDataSource(cfg.DataBase)
		if err != nil {
			panic(err)
		}

		db, err := gormDB.DB()
		if err != nil {
			panic(err)
		}

		if err := goose.Up(db, "."); err != nil {
			panic(err)
		}
	},
}
