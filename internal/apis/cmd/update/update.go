package update

import (
	"fmt"
	"os"
	"strings"

	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	cfg *config.Config

	Cmd = &cobra.Command{
		Use:       "update [RESOURCE]",
		Short:     "Update one resource",
		ValidArgs: []string{"tasks"},
		Args:      cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			// todo: 2021-02-22|23:46|doggy|refactor me
			switch args[1] {
			case "status":
				if status, ok := pb.Status_value[strings.ToUpper(args[2])]; !ok {
					fmt.Printf("status parse error %v\n", args[2])
				} else {
					// todo: 2021-02-22|23:57|doggy|implement me
					fmt.Printf("update %v resource %v to %v", args[0], args[1], status)
				}
				break
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".lobster")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	_ = viper.ReadInConfig()

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
}
