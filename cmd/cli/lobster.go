package main

import (
	"os"

	"github.com/blackhorseya/lobster/internal/cmd"
)

func main() {
	if err := cmd.NewCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
