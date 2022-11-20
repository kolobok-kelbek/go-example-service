//go:build tools
// +build tools

// this file need for add some packages to vendor directory
// it's not a joke, see https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/google/wire/cmd/wire"
)
