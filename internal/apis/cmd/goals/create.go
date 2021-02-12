package goals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create TITLE",
	Short: "Create a objective",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uri := fmt.Sprintf("%v/v1/goals", cfg.API.EndPoint)
		data, _ := json.Marshal(&okr.Objective{Title: args[0]})
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

		var obj *okr.Objective
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader(header)
		table.Append(obj.ToLine())
		table.Render()
	},
}

func init() {
	Cmd.AddCommand(createCmd)
}
