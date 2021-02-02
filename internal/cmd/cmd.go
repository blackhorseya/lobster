package cmd

import (
	"github.com/blackhorseya/lobster/internal/cmd/config"
	"github.com/blackhorseya/lobster/internal/cmd/tasks"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use: "lobster",
		Long: `
██╗      ██████╗ ██████╗ ███████╗████████╗███████╗██████╗ 
██║     ██╔═══██╗██╔══██╗██╔════╝╚══██╔══╝██╔════╝██╔══██╗
██║     ██║   ██║██████╔╝███████╗   ██║   █████╗  ██████╔╝
██║     ██║   ██║██╔══██╗╚════██║   ██║   ██╔══╝  ██╔══██╗
███████╗╚██████╔╝██████╔╝███████║   ██║   ███████╗██║  ██║
╚══════╝ ╚═════╝ ╚═════╝ ╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝

Lobster is a tool which integration todo list, OKRs, sprint board, pomodoro and report etc. functional.
`,
		Version: "1.0.0",
	}
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.AddCommand(config.Cmd)
	rootCmd.AddCommand(tasks.Cmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lobster.yaml)")
}

// NewInjector serve caller to create an injector
func NewCommand() *cobra.Command {
	return rootCmd
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewCommand)
