package config

import (
	"embed"
	"github.com/kolobok-kelbek/cong"
)

func Load(snapshot embed.FS) Config {
	loader := cong.NewLoader[Config]()
	cfg, err := loader.LoadFromEmbedFS("Tomato", snapshot, cong.YamlExt)
	if err != nil {
		panic(err)
	}

	return *cfg
}
