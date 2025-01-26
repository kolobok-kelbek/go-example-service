package config

import (
	"fmt"
	"time"
)

type DataBase struct {
	Dialect              string
	DSN                  string
	User                 string
	Password             string
	Host                 string
	Port                 int
	Database             string
	MaxLifeTime          int
	PreferSimpleProtocol bool
	InsertBatchSize      int
	SlowThreshold        int
}

func (c *DataBase) GetTimeout() time.Duration {
	return time.Duration(c.MaxLifeTime) * time.Second
}

func (c *DataBase) Dsn() string {
	switch c.Dialect {
	case "postgres":
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.User, c.Password, c.Host, c.Port, c.Database)
	}

	return ""
}
