package config

import (
	db "github.com/kolobok-kelbek/tomato/internal/db/config"
)

type Config struct {
	Server   Server
	DataBase db.DataBase
}
