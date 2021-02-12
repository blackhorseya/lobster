package goals

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/olekukonko/tablewriter"
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

		var goals []*okr.Objective
		err = json.Unmarshal(body, &goals)
		if err != nil {
			fmt.Println(err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader(header)
		for _, g := range goals {
			table.Append(g.ToLine())
		}
		table.Render()
	},
}

func init() {
	Cmd.AddCommand(listCmd)

	listCmd.Flags().Int("page", 1, "list objectives which page")
	listCmd.Flags().Int("size", 1, "list objectives which size")
}
