package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
	"github.com/spf13/cobra"
)

var (
	tasksCmd = &cobra.Command{
		Use: "tasks",
	}

	tasksListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run:   list,
	}

	tasksCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a task",
		Run:   create,
	}
)

func init() {
	rootCmd.AddCommand(tasksCmd)

	tasksCmd.AddCommand(tasksListCmd)
	tasksListCmd.Flags().Int("page", 1, "page")
	tasksListCmd.Flags().Int("size", 10, "size")

	tasksCmd.AddCommand(tasksCreateCmd)
	tasksCreateCmd.Flags().String("title", "", "title")
}

func list(cmd *cobra.Command, args []string) {
	page, _ := cmd.Flags().GetInt("page")
	size, _ := cmd.Flags().GetInt("size")

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
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func create(cmd *cobra.Command, args []string) {
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
}
