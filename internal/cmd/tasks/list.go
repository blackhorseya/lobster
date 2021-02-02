package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
	"github.com/spf13/cobra"
)

const (
	format = "%-36s\t%-20s\t%-6v\t%-9v"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
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
		defer resp.Body.Close()

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

		ret := []string{fmt.Sprintf(format, "ID", "Title", "Status", "Create At")}
		for _, t := range tasks {
			ret = append(ret, t.ToLineByFormat(format))
		}

		fmt.Println(strings.Join(ret, "\n"))
	},
}

func init() {
	Cmd.AddCommand(listCmd)

	listCmd.Flags().Int("page", 1, "list tasks which page")
	listCmd.Flags().Int("size", 10, "list tasks which size")
}
