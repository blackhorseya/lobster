package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task by id",
	Long:  "lobster tasks update ID [flags]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return er.ErrInvalidID
		}

		if _, err := uuid.Parse(args[0]); err != nil {
			return er.ErrInvalidID
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		completed, _ := cmd.Flags().GetBool("status")
		uri := fmt.Sprintf("%v/v1/tasks/%v", cfg.API.EndPoint, args[0])
		data, _ := json.Marshal(&todo.Task{
			ID:        args[0],
			Title:     title,
			Completed: completed,
		})
		req, err := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(data))
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
	Cmd.AddCommand(updateCmd)

	updateCmd.Flags().String("title", "", "title of task")
	updateCmd.Flags().Bool("status", false, "completed of task")
}
