package config

import (
	db "github.com/kolobok-kelbek/tomato/pkg/db/config"
)

type Config struct {
	Server   Server
	DataBase db.DataBase
}
