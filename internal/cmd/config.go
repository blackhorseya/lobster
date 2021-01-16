package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use: "config",
	}

	configViewCmd = &cobra.Command{
		Use:   "view",
		Short: "Show configuration",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cfg.String())
		},
	}
)

func init() {
	configCmd.AddCommand(configViewCmd)

	rootCmd.AddCommand(configCmd)
}
