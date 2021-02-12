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

var getCmd = &cobra.Command{
	Use:   "get ID",
	Short: "Get objective by ID",
	Long:  "lobster objectives get ID [flags]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uri := fmt.Sprintf("%v/v1/goals/%v", cfg.API.EndPoint, args[0])
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

		var obj *okr.Objective
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err)
			return
		}

		ret := []string{header, obj.ToLineByFormat(format)}
		fmt.Println(strings.Join(ret, "\n"))
	},
}

func init() {
	Cmd.AddCommand(getCmd)
}
