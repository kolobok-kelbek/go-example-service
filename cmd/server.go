package cmd

import (
	configDir "github.com/kolobok-kelbek/tomato/config"
	"github.com/kolobok-kelbek/tomato/internal/app"
	"github.com/kolobok-kelbek/tomato/internal/config"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "up server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load(configDir.Snapshot, config.ProdEnv)
		if err != nil {
			panic(err)
		}

		di, err := app.NewDIContainer(cfg)
		if err != nil {
			panic(err)
		}

		err = di.App().Run()
		if err != nil {
			panic(err)
		}
	},
}
