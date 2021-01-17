package cmd

import (
	"fmt"
	"os"

	"github.com/blackhorseya/lobster/internal/cmd/config"
	C "github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/spf13/cobra"
)

var (
	cfgFile string

	cfg *C.Config

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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.AddCommand(config.Cmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lobster.yaml)")
}
