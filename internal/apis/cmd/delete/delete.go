package delete

import (
	"fmt"
	"net/http"
	"os"

	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	cfg *config.Config

	// Cmd is a root command for delete
	Cmd = &cobra.Command{
		Use:       "delete [RESOURCE]",
		Short:     "Delete one resource",
		ValidArgs: []string{"tasks", "results", "goals"},
		Args:      cobra.ExactArgs(2),
		Aliases:   []string{"del", "remove"},
		Run: func(cmd *cobra.Command, args []string) {
			resource := args[0]
			id := args[1]

			uri := fmt.Sprintf("%v/v1/%v/%v", cfg.API.EndPoint, resource, id)
			req, err := http.NewRequest(http.MethodDelete, uri, nil)
			if err != nil {
				fmt.Println(err)
				return
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusNoContent {
				fmt.Printf("Delete %v ID: %v is success\n", resource, id)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	Cmd.Flags().Int("page", 1, "list resource which page")
	Cmd.Flags().Int("size", 10, "list resource which size")
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
