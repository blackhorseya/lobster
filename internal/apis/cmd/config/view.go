package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cfg.String())
	},
}
