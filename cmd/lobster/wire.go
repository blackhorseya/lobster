// +build wireinject

package main

import (
	"github.com/blackhorseya/lobster/internal/cmd"
	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

var providerSet = wire.NewSet(
	cmd.ProviderSet,
	config.ProviderSet,
)

// CreateCommand serve caller to create a command
func CreateCommand() (*cobra.Command, error) {
	panic(wire.Build(providerSet))
}
