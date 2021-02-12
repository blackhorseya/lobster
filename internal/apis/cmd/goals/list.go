package goals

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all objectives",
	Run: func(cmd *cobra.Command, args []string) {
		page, _ := cmd.Flags().GetInt("page")
		size, _ := cmd.Flags().GetInt("size")

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

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(body)
			return
		}

		var objs []*okr.Objective
		err = json.Unmarshal(body, &objs)
		if err != nil {
			fmt.Println(err)
			return
		}

		ret := []string{header}
		for _, obj := range objs {
			ret = append(ret, obj.ToLineByFormat(format))
		}

		fmt.Println(strings.Join(ret, "\n"))
	},
}

func init() {
	Cmd.AddCommand(listCmd)

	listCmd.Flags().Int("page", 1, "list objectives which page")
	listCmd.Flags().Int("size", 1, "list objectives which size")
}
