//go:build tools
// +build tools

package main

import (
	_ "github.com/air-verse/air"
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/go-task/task/v3/cmd/task"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/pressly/goose/v3/cmd/goose"
)
