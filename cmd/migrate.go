package cmd

import (
	configFiles "github.com/kolobok-kelbek/tomato/config"
	"github.com/kolobok-kelbek/tomato/internal/config"
	"github.com/kolobok-kelbek/tomato/migrations"
	database "github.com/kolobok-kelbek/tomato/pkg/db"
	"github.com/kolobok-kelbek/tomato/pkg/db/dialect"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		goose.SetBaseFS(migrations.Snapshot)

		cfg := config.Load(configFiles.Snapshot)

		if err := goose.SetDialect(cfg.DataBase.Dialect); err != nil {
			panic(err)
		}

		dialectManager := dialect.NewManager()
		gormDB := database.GetDataSource(cfg.DataBase, dialectManager)

		db, err := gormDB.DB()
		if err != nil {
			panic(err)
		}

		if err := goose.Up(db, "."); err != nil {
			panic(err)
		}
	},
}
