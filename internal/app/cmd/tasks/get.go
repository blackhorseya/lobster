package tasks

import (
	"fmt"
	"io/ioutil"
	"net/http"

	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get task by ID",
	Long: "lobster tasks get ID [flags]",
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
		uri := fmt.Sprintf("%v/v1/tasks/%v", cfg.API.EndPoint, args[0])
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
	Cmd.AddCommand(getCmd)
}
