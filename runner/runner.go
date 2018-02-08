package runner

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"

	"log"

	"crypto/tls"
	"crypto/x509"
	"git.elenet.me/yuelong.huang/pansible/client"
	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/queue"
	"github.com/Sirupsen/logrus"
	"github.com/franela/goreq"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Runner struct {
	Token        string `json:"-"`
	Master       string `json:"master"`
	Task         string `json:"task"`
	SSHKeyServer string
	Port         int
	ListenServer *echo.Echo
	*models.Run
	Playbook  *models.Playbook
	done      chan bool
	callbackQ *queue.Queue
}

func New(token string) *Runner {
	return &Runner{
		Token:        token,
		Master:       viper.GetString("MASTER"), //claims.Master,
		Task:         viper.GetString("RUN_ID"), //claims.Task,
		SSHKeyServer: viper.GetString("SSHKEY_SERVER"),
		//JobID:     viper.GetString("JOB_ID"),
		Port:      viper.GetInt("RUNNER_PORT"),
		done:      make(chan bool),
		callbackQ: queue.NewQueue(),
	}
}

func (r *Runner) getSSHKey() error {
	cert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		return errors.Wrap(err, "Couldn't load file")
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)

	conf := &tls.Config{
		RootCAs: certPool,
	}

	tr := &http.Transport{
		TLSClientConfig: conf,
	}
	c := &http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s/runner/ssh-key", r.SSHKeyServer), nil)
	if err != nil {
		return errors.Wrap(err, "failed to new sshkey request")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.Token))
	resp, err := c.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send sshkey request to server")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return errors.Errorf("status code %d", resp.StatusCode)
	}

	if err := os.MkdirAll("./tmp/pansible", os.ModePerm); err != nil {
		return errors.Wrap(err, "failed to ensure tmp dir")
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	return errors.Wrap(
		ioutil.WriteFile("./tmp/pansible/id", content, 0600),
		"failed to write ssh keyfile",
	)
}

func (r *Runner) getPlaybook(client *client.Client) error {
	if err := os.MkdirAll("./tmp/pansible", os.ModePerm); err != nil {
		return errors.Wrap(err, "failed to ensure tmp dir")
	}

	playbook, err := client.GetPlaybookByJobID(r.JobID)
	if err != nil {
		return errors.Wrap(err, "failed to get playbook")
	}
	r.Playbook = playbook

	playbookDir := "./tmp/pansible/playbook_" + r.JobID
	if _, err := os.Stat(playbookDir); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", playbook.GitRepo, "./tmp/pansible/playbook_"+r.JobID)
		cmd.Env = append(os.Environ(), "GIT_SSH_COMMAND=ssh -o StrictHostKeyChecking=no -i ./tmp/pansible/id")
		_, err := cmd.CombinedOutput()
		if err != nil {
			return errors.Wrap(err, "failed to clone playbook repo")
		}
	}
	return nil
}

func (r *Runner) getRun(client *client.Client) error {

	var err error
	r.Run, err = client.GetRun(r.Task)
	if err != nil {
		return errors.Wrap(err, "client failed to get run by id")
	}
	return nil
}

func (r *Runner) Setup() error {
	c := client.New(r.Master, viper.GetString("token"))

	if err := r.getSSHKey(); err != nil {
		return errors.Wrap(err, "failed to get ssh key")
	}

	if err := r.getRun(c); err != nil {
		return errors.Wrap(err, "failed to get run information")
	}

	if err := r.getPlaybook(c); err != nil {
		return errors.Wrap(err, "failed to get playbook")
	}

	return nil
}

func (r *Runner) StartServer() {
	app := echo.New()
	r.ListenServer = app

	app.POST("/events", r.ProcessCallback)

	logrus.Fatal(app.Start(fmt.Sprintf(":%d", r.Port)))

}

func (r *Runner) ProcessCallback(ctx echo.Context) error {
	var body echo.Map
	log.Printf("runner callback received\n")
	if err := ctx.Bind(&body); err != nil {
		log.Printf("runner callback bind failed\n")
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "runner failed to parse callback",
		})
	}
	log.Printf("runner callback received => %v\n", body)

	r.callbackQ.Enqueue(body)
	return nil
}

func (r *Runner) MasterFeedback() {
	for c := range r.callbackQ.Dequeue() {
		log.Printf("dequeued c => %v\n", c)
		req := goreq.Request{
			Method: http.MethodPost,
			Uri:    fmt.Sprintf("http://%s/runner/events", r.Master),
			Body:   c,
		}.WithHeader("Authorization", fmt.Sprintf("Bearer %s", r.Token))

		for {
			resp, err := req.Do()
			if err != nil {
				logrus.WithError(err).Errorln("failed to send callback")
				time.Sleep(time.Second)
				continue
			}

			if resp.StatusCode < 400 && resp.StatusCode >= 200 {
				break
			}

			time.Sleep(time.Second)
		}
	}
	log.Println("feedServer ended")
}

func (r *Runner) Exec() error {
	timeout := 10 * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	args := []string{
		"--key-file", "./tmp/pansible/id",
		//"-vvvv",
		"-i", os.Args[0],
		fmt.Sprintf("tmp/pansible/playbook_%s/%s", r.JobID, r.Playbook.Entry),
	}

	if r.Limit != "" {
		if err := ioutil.WriteFile(r.UUID+".limit", []byte(r.Limit), 0666); err != nil {
			return errors.Wrap(err, "failed to write limit file")
		}
		args = append(args, "--limit", fmt.Sprintf("@%s.limit", r.UUID))
	}

	cmd := exec.CommandContext(ctx, "ansible-playbook", args...)

	cmd.Env = []string{
		"ANSIBLE_HOST_KEY_CHECKING=False",
		fmt.Sprintf("PANSIBLE_RUN_ID=%s", r.Task),
		fmt.Sprintf("PANSIBLE_JOB_ID=%s", r.JobID),
		fmt.Sprintf("PANSIBLE_TOKEN=%s", r.Token),
		fmt.Sprintf("PANSIBLE_RUNNER=localhost:%d", r.Port),
		fmt.Sprintf("PANSIBLE_MASTER=%s", r.Master),
	}

	output, err := cmd.CombinedOutput()

	fmt.Println(string(output))

	var reason string
	if err != nil {
		reason = fmt.Sprintf("failed: %s", err.Error())
	} else {
		reason = "success"
	}

	r.callbackQ.Enqueue(echo.Map{
		"_pansible_task": r.Task,
		"_pansible_job":  r.JobID,
		"type":           "task_finish",
		"reason":         reason,
	})
	r.callbackQ.Stop()

	if err := os.RemoveAll(fmt.Sprintf("%s.limit", r.UUID)); err != nil {
		logrus.WithError(err).Println("runner failed to remove limit file ")
	}

	return errors.Wrap(err, "command failed")
}
