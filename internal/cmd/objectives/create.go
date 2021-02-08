package objectives

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a objective",
	Long:  "lobster objectives create TITLE [flags]",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uri := fmt.Sprintf("%v/v1/objectives", cfg.API.EndPoint)
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

		ret := []string{header, obj.ToLineByFormat(format)}

		fmt.Println(strings.Join(ret, "\n"))
	},
}

func init() {
	Cmd.AddCommand(createCmd)
}
