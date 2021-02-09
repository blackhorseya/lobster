package main

import (
	"os"

	cmd2 "github.com/blackhorseya/lobster/internal/apis/cmd"
)

func main() {
	if err := cmd2.NewCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
