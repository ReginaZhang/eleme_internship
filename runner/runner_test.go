package runner

import (
	"encoding/json"
	"fmt"
	"git.elenet.me/yuelong.huang/pansible/api"
	"git.elenet.me/yuelong.huang/pansible/client"
	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/tests"
	"github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
	"github.com/franela/goreq"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestGetPlaybook(t *testing.T) {
	s := tests.NewTestServer(t)
	a := assert.New(t)

	//taskID := uuid.NewV4().String()
	taskID := "10bff853-3c41-4c19-af9c-45b694794fde"
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, api.Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
		"10bff853-3c41-4c19-af9c-45b694794fde",
		"localhost:5757",
		"test",
	})

	token, err := claim.SignedString([]byte("secret"))
	a.Nil(err)

	client, err := client.New(token)
	a.Nil(err)

	fmt.Println("parse succeeded")

	playbook := models.Playbook{
		Name:    "test",
		GitRepo: "git@some_repo",
		Entry:   "some_playbook.yml",
	}

	err = playbook.Insert(s.DB)
	a.Nil(err)

	play := models.Job{
		InventoryID: 1,
		PlaybookID:  1,
		Env:         "for testing's sake",
		UUID:        uuid.NewV4().String(),
	}

	err = play.Insert(s.DB)
	a.Nil(err)

	run := models.Run{
		JobID: "1",
		Env:   "for testing's sake",
		UUID:  taskID,
	}

	err = run.Insert(s.DB)
	a.Nil(err)

	r, err := New(token)
	a.Nil(err)

	err = r.getPlaybook(client)
	a.Nil(err)

	_, err = os.Stat("/tmp/pansible/playbook")
	a.Nil(err)

}

func TestRunner(t *testing.T) {

	s := tests.NewTestServer(t)
	a := assert.New(t)

	task := models.Run{
		JobID: "123",
		Env:   "for testing's sake",
		UUID:  "10bff853-3c41-4c19-af9c-45b694794fde",
	}

	err := task.Insert(s.DB)
	a.Nil(err)

	job := models.Job{
		InventoryID: 1,
		PlaybookID:  1,
		UUID:        "123",
		Env:         "for testing's sake",
	}

	err = job.Insert(s.DB)
	a.Nil(err)

	pb := models.Playbook{
		Name:    "test",
		GitRepo: "git@some_repo",
		Entry:   "some_playbook.yml",
	}
	err = pb.Insert(s.DB)
	a.Nil(err)

	content, err := ioutil.ReadFile("/tmp/test.yml")
	a.Nil(err)

	fmt.Println("host string => ", string(content))
	inv := models.Inventory{
		Env:     "alpha",
		Version: "0.0.1",
		Hosts:   string(content),
		Vars:    "---",
	}
	err = inv.Insert(s.DB)
	a.Nil(err)

	token := getToken("10bff853-3c41-4c19-af9c-45b694794fde")
	viper.Set("token", token)

	r, err := New(token)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to create runner")
	}

	if err := r.Setup(); err != nil {
		logrus.WithError(err).Fatal("Failed to setup")
	}

	// err := Run()
	// fmt.Println(err)
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
		Uri:     "http://localhost:5757/login",
		Body:    body,
	}

	fmt.Println(string(body))

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

	fmt.Println(resp.Token)

	return resp.Token
}
