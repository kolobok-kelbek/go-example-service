// Package static is entry point for load all static files in memory
package config

import "embed"

// Snapshot is copy of static files
//
//go:embed *.yaml
var Snapshot embed.FS
