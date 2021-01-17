package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use: "create",
	Short: "Create a task",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		uri := fmt.Sprintf("%v/v1/tasks", cfg.API.EndPoint)
		data, _ := json.Marshal(&todo.Task{Title: title})
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

		fmt.Println(string(body))
	},
}

func init() {
	createCmd.Flags().String("title", "", "title of task")
}
