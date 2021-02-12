package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a task",
	Long:  "lobster tasks create TITLE [flags]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return er.ErrEmptyTitle
		}

		if len(args[0]) == 0 {
			return er.ErrEmptyTitle
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		uri := fmt.Sprintf("%v/v1/tasks", cfg.API.EndPoint)
		data, _ := json.Marshal(&todo.Task{Title: args[0]})
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
		defer func() {
			_ = resp.Body.Close()
		}()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var task *todo.Task
		err = json.Unmarshal(body, &task)
		if err != nil {
			fmt.Println(err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader(header)
		table.Append(task.ToLine())
		table.Render()
	},
}

func init() {
	Cmd.AddCommand(createCmd)
}
