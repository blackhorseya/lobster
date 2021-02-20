package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	cfg *config.Config

	Cmd = &cobra.Command{
		Use:       "create [RESOURCE] [TITLE]",
		Short:     "Create one resource",
		ValidArgs: []string{"tasks", "results", "goals"},
		Args:      cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			// todo: 2021-02-15|19:46|doggy|refactor me
			uri := fmt.Sprintf("%v/v1/%v", cfg.API.EndPoint, args[0])

			switch args[0] {
			case "tasks":
				data, _ := json.Marshal(&todo.Task{Title: args[1]})
				req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
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

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				var ret *todo.Task
				err = json.Unmarshal(body, &ret)
				if err != nil {
					fmt.Println(err)
					return
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
				table.SetHeader([]string{"ID", "Title", "Status", "Create At"})
				table.Append(ret.ToLine())
				table.Render()

				break
			case "goals":
				data, _ := json.Marshal(&okr.Objective{Title: args[1]})
				req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
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

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				var ret *okr.Objective
				err = json.Unmarshal(body, &ret)
				if err != nil {
					fmt.Println(err)
					return
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
				table.SetHeader([]string{"ID", "Title", "Start At", "End At", "Create At"})
				table.Append(ret.ToLine())
				table.Render()

				break
			case "results":
				if len(cfg.Context.Goal) == 0 {
					fmt.Println("missing context.goal in .lobster.yaml")
					return
				}

				_, err := uuid.Parse(cfg.Context.Goal)
				if err != nil {
					fmt.Println(err)
					return
				}

				data, _ := json.Marshal(&okr.KeyResult{Title: args[1], GoalID: cfg.Context.Goal})
				req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
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

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				var ret *okr.KeyResult
				err = json.Unmarshal(body, &ret)
				if err != nil {
					fmt.Println(err)
					return
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
				table.SetHeader([]string{"ID", "Goad ID", "Title", "Target", "Actual", "Create At"})
				table.Append(ret.ToLine())
				table.Render()

				break
			}

			return
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
