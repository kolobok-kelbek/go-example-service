package config

import (
	"embed"
	"github.com/kolobok-kelbek/cong"
)

const (
	ProdEnv = "prod"
	TestEnv = "test"

	prodPath = "prod.yaml"
	testPath = "test.yaml"
)

func Load(snapshot embed.FS, env string) (*Config, error) {
	path := prodPath
	if env == TestEnv {
		path = testPath
	}

	loader := cong.NewLoader[Config]()
	cfg, err := loader.LoadFromEmbedFSByPath("Tomato", snapshot, path, cong.YamlExt)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
