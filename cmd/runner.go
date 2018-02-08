// +build !server_only

package cmd

import (
	"fmt"

	"git.elenet.me/yuelong.huang/pansible/runner"

	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/franela/goreq"
	"github.com/labstack/echo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//TODO: errors needs to log to files not to standard out!
// runnerCmd represents the runner command
var runnerCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile("runner.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		exitOnError(err, "Failed to open log file")

		log.SetOutput(f)
		logrus.SetOutput(f)

		// TODO: Work your own magic here
		viper.Set("RUN_ID", args[0])
		viper.Set("JOB_ID", args[1])
		token := getToken(args[0])
		viper.Set("token", token)

		r := runner.New(token)

		logrus.WithFields(logrus.Fields{
			"task":   r.Task,
			"master": r.Master,
			"port":   r.Port,
		}).Infoln("Start job")

		exitOnError(r.Setup(), "Failed to setup")

		go r.Exec()

		go r.StartServer()

		r.MasterFeedback()

		log.Println("runner ended")
	},
}

func init() {
	RootCmd.AddCommand(runnerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runnerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runnerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func getToken(uuid string) string {
	bodyMap := map[string]string{
		"run_uuid": uuid,
	}
	body, err := json.Marshal(bodyMap)
	if err != nil {
		logrus.WithError(err).Fatal("Json failed to marshal request body to get token")
	}
	request := goreq.Request{
		Timeout: time.Second,
		Method:  http.MethodPost,
		Uri:     fmt.Sprintf("http://%s/runnerticket", viper.GetString("MASTER")),
		Body:    body,
	}

	request.AddHeader(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response, err := request.Do()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get token")
	}
	defer response.Body.Close()

	var resp struct {
		Token string `json:"token"`
	}
	if err := response.Body.FromJsonTo(&resp); err != nil {
		logrus.WithError(err).Fatal("Failed to parse login response body")
	}

	return resp.Token
}
