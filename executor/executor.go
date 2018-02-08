package executor

import (
	"fmt"
	"git.elenet.me/yuelong.huang/pansible/models"
	"github.com/Sirupsen/logrus"
	"github.com/franela/goreq"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Executor interface {
	Run(master, sshkeyServer string, runs []*models.Run) error
}

type LocalExecutor struct{}

type ApposExecutor struct{}

func NewLocal() *LocalExecutor {
	return &LocalExecutor{}
}

func NewAppos() *ApposExecutor {
	return &ApposExecutor{}
}

func (e *LocalExecutor) Run(master, sshkeyServer string, runs []*models.Run) error {
	var err error
	for _, run := range runs {
		go func() {
			_, err = filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				logrus.WithError(err).Printf("Failed to find executable path")
			}
			cmd := exec.Command(os.Args[0], "run", run.UUID, run.JobID)

			cmd.Env = append(os.Environ(),
				fmt.Sprintf("PANSIBLE_MASTER=%s", master),
				fmt.Sprintf("PANSIBLE_RUNNER_PORT=%s", "12345"),
				fmt.Sprintf("PANSIBLE_SSHKEY_SERVER=%s", sshkeyServer))
			output, err := cmd.CombinedOutput()
			if err != nil {
				logrus.WithError(err).Printf("Failed to start pansible run")
			}
			fmt.Println(string(output))
		}()
	}

	return err
}

func (e *ApposExecutor) Run(master, sshkeyServer string, runs []*models.Run) error {
	for _, run := range runs {
		go func() {
			env := append(os.Environ(),
				fmt.Sprintf("PANSIBLE_MASTER=%s", master),
				fmt.Sprintf("PANSIBLE_RUNNER_PORT=%s", "12345"),
				fmt.Sprintf("PANSIBLE_SSHKEY_SERVER=%s", sshkeyServer))
			var params []map[string]string
			for _, e := range env {
				params = append(params, map[string]string{
					"key":   "env",
					"value": e,
				})
			}
			resp, err := goreq.Request{
				Method:      http.MethodPost,
				Uri:         "http://some_remote_service",
				ContentType: "application/json",
				Body: map[string]interface{}{
					"some_key": "some_value",
					"params":   params,
				},
			}.Do()
			if err != nil {
				logrus.WithError(err).Printf("Failed to start pansible run using Appos executor")
				return
			}
			var respBody echo.Map
			resp.Body.FromJsonTo(respBody)
		}()
	}

	return nil
}
