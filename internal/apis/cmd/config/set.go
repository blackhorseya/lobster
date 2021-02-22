package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setCmd = &cobra.Command{
	Use: "set [FIELD] [VALUE]",
	Short: "Set configuration",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		field := args[0]
		value := args[1]

		viper.Set(field, value)
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
