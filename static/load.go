// Package static is entry point for load all static files in memory
package static

import "embed"

// Snapshot is copy of static files
//go:embed *
var Snapshot embed.FS
