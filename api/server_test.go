package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/tests"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func newReq(method, url string, body interface{}) *http.Request {
	b, err := json.Marshal(body)

	if err != nil {
		return nil
	}

	req := httptest.NewRequest(method, url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	return req
}

func createApp(s *Server, app models.App) error {
	e := echo.New()
	req := newReq(echo.PUT, "/api/v1/app", app)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	return s.CreateApp(ctx)
}

func TestCreateApp(t *testing.T) {
	t.Skip()
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	req := newReq(echo.PUT, "/api/v1/app", models.App{
		Appid: "tools.test",
	})

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	err := s.CreateApp(ctx)
	a.Nil(err)
}

func TestDeployApp(t *testing.T) {
	t.Skip()
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	tests.FailOnError(t, createApp(&s, models.App{
		Appid:    "tools.test",
		TaskRepo: "some_repo",
		Playbook: "some_playbook.yml",
	}))

	invContent := `some_yaml_content
`

	varBytes, _ := json.Marshal(echo.Map{
		"a": 1,
	})

	tests.FailOnError(t, createInventory(&s, models.Inventory{
		Env:     "alpha",
		Version: "0.0.1",
		Hosts:   invContent,
		Vars:    string(varBytes),
	}))

	req := newReq(echo.POST, "/api/v1/deploy", echo.Map{
		"appid":          "tools.test",
		"env":            "alpha",
		"config_version": "0.0.1",
	})

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	err := s.DeployApp(ctx)
	a.Nil(err)
}
