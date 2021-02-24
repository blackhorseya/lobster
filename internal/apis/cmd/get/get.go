package get

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	cfg *config.Config

	Cmd = &cobra.Command{
		Use:       "get [RESOURCE]",
		Short:     "Display one resource",
		ValidArgs: []string{"tasks", "results", "goals"},
		Args:      cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			page, _ := cmd.Flags().GetInt("page")
			size, _ := cmd.Flags().GetInt("size")

			// todo: 2021-02-15|14:46|doggy|refactor it

			if len(args) == 1 {
				uri := fmt.Sprintf("%v/v1/%v?page=%v&size=%v", cfg.API.EndPoint, args[0], page, size)
				req, err := http.NewRequest(http.MethodGet, uri, nil)
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

				if resp.StatusCode == http.StatusNotFound {
					fmt.Println("No resources found")
					return
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				switch args[0] {
				case "tasks":
					var data []*pb.Task
					err = json.Unmarshal(body, &data)
					if err != nil {
						fmt.Println(err)
						return
					}

					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
					table.SetHeader([]string{"ID", "Result ID", "Title", "Status", "Create At"})
					for _, t := range data {
						table.Append(t.ToLine())
					}
					table.Render()

					break
				case "goals":
					var data []*pb.Goal
					err = json.Unmarshal(body, &data)
					if err != nil {
						fmt.Println(err)
						return
					}

					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
					table.SetHeader([]string{"ID", "Title", "Start At", "End At", "Create At"})
					for _, g := range data {
						table.Append(g.ToLine())
					}
					table.Render()

					break
				case "results":
					var data []*pb.Result
					err = json.Unmarshal(body, &data)
					if err != nil {
						fmt.Println(err)
						return
					}

					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
					table.SetHeader([]string{"ID", "Goal ID", "Title", "Target", "Actual", "Progress", "Create At"})
					for _, t := range data {
						table.Append(t.ToLine())
					}
					table.Render()

					break
				}
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
