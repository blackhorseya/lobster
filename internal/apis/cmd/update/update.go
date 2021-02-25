package update

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

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
		Use:       "update [RESOURCE]",
		Short:     "Update one resource",
		ValidArgs: []string{"tasks", "goals", "results"},
		Args:      cobra.MinimumNArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			resource := args[0]
			id := args[1]
			field := args[2]
			value := args[3]

			// todo: 2021-02-22|23:46|doggy|refactor me
			uri := fmt.Sprintf("%v/v1/%v/%v/%v", cfg.API.EndPoint, resource, id, field)

			switch field {
			case "status":
				if status, ok := pb.Status_value[strings.ToUpper(value)]; !ok {
					fmt.Printf("status parse error %v\n", value)
				} else {
					data, _ := json.Marshal(&pb.Task{Status: pb.Status(status)})
					req, err := http.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(data))
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

					var task *pb.Task
					err = json.Unmarshal(body, &task)
					if err != nil {
						fmt.Println(err)
						return
					}

					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
					table.SetHeader([]string{"ID", "Result ID", "Title", "Status", "Create At"})
					table.Append(task.ToLine())
					table.Render()
				}
				break
			case "title":
				switch resource {
				case "tasks":
					data, _ := json.Marshal(&pb.Task{Title: value})
					req, err := http.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(data))
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

					var task *pb.Task
					err = json.Unmarshal(body, &task)
					if err != nil {
						fmt.Println(err)
						return
					}

					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
					table.SetHeader([]string{"ID", "Result ID", "Title", "Status", "Create At"})
					table.Append(task.ToLine())
					table.Render()

					break
				case "goals":
					data, _ := json.Marshal(&pb.Goal{Title: value})
					req, err := http.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(data))
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

					var task *pb.Goal
					err = json.Unmarshal(body, &task)
					if err != nil {
						fmt.Println(err)
						return
					}

					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
					table.SetHeader([]string{"ID", "Title", "Start At", "End At", "Create At"})
					table.Append(task.ToLine())
					table.Render()

					break
				case "results":
					data, _ := json.Marshal(&pb.Result{Title: value})
					req, err := http.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(data))
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

					var ret *pb.Result
					err = json.Unmarshal(body, &ret)
					if err != nil {
						fmt.Println(err)
						return
					}

					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
					table.SetHeader([]string{"ID", "Goal ID", "Title", "Target", "Actual", "Progress", "Create At"})
					table.Append(ret.ToLine())
					table.Render()

					break
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
