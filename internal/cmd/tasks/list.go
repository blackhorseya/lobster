package tasks

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
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
	listCmd.Flags().Int("page", 1, "list tasks which page")
	listCmd.Flags().Int("size", 10, "list tasks which size")
}
