package get

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
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
		Args:      cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			page, _ := cmd.Flags().GetInt("page")
			size, _ := cmd.Flags().GetInt("size")

			switch args[0] {
			case "tasks":
				// todo: 2021-02-15|14:46|doggy|refactor it
				uri := fmt.Sprintf("%v/v1/tasks?page=%v&size=%v", cfg.API.EndPoint, page, size)
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
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				var tasks []*todo.Task
				err = json.Unmarshal(body, &tasks)
				if err != nil {
					fmt.Println(err)
					return
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
				table.SetHeader([]string{"ID", "Title", "Status", "Create At"})
				for _, t := range tasks {
					table.Append(t.ToLine())
				}
				table.Render()

				break
			case "goals":
				// todo: 2021-02-15|14:52|doggy|refactor it
				uri := fmt.Sprintf("%v/v1/goals?page=%v&size=%v", cfg.API.EndPoint, page, size)
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
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(body)
					return
				}

				var goals []*okr.Objective
				err = json.Unmarshal(body, &goals)
				if err != nil {
					fmt.Println(err)
					return
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
				table.SetHeader([]string{"ID", "Title", "Start At", "End At", "Create At"})
				for _, g := range goals {
					table.Append(g.ToLine())
				}
				table.Render()

				break
			case "results":
				// todo: 2021-02-15|14:52|doggy|refactor it
				uri := fmt.Sprintf("%v/v1/krs?page=%v&size=%v", cfg.API.EndPoint, page, size)
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
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				var tasks []*okr.KeyResult
				err = json.Unmarshal(body, &tasks)
				if err != nil {
					fmt.Println(err)
					return
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
				table.SetHeader([]string{"ID", "Goad ID", "Title", "Target", "Actual", "Create At"})
				for _, t := range tasks {
					table.Append(t.ToLine())
				}
				table.Render()

				break
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
