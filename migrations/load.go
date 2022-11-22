package migrations

import "embed"

// Snapshot is copy of migration files
//
//go:embed *.sql
var Snapshot embed.FS
