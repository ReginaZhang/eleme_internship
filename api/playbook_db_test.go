package api

import (
	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/tests"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func createPlaybook(s *Server, pb models.Playbook) error {
	e := echo.New()
	req := newReq(echo.POST, "/api/v1/playbook", pb)

	ctx := e.NewContext(req, httptest.NewRecorder())

	return s.CreatePlaybook(ctx)
}

func TestListPlaybooks(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	tests.FailOnError(t, createPlaybook(&s, models.Playbook{
		Name:    "test",
		GitRepo: "https://github.com/ansible/ansible-examples/tree/master/",
		Entry:   "tomcat-standalone/roles/tomcat/tasks/main.yml",
	}))

	req := newReq(echo.GET, "/api/v1/playbook", nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	err := s.ListPlaybooks(ctx)

	a.Equal(http.StatusOK, ctx.Response().Status)
	a.Nil(err)
}

func TestDeletePlaybook(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	tests.FailOnError(t, createPlaybook(&s, models.Playbook{
		Name:    "test",
		GitRepo: "git@some_repo",
		Entry:   "some_playbook.yml",
	}))

	tests.FailOnError(t, createJob(&s, models.Job{
		InventoryID: 1,
		PlaybookID:  1,
		Env:         "for testing's sake",
	}))

	tests.FailOnError(t, createRun(&s, models.Run{
		JobID: "1",
		Env:   "for testing's sake",
	}))

	pb, err := models.Playbooks(s.db).One()

	a.Nil(err)

	req := newReq(echo.DELETE, "/", nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	ctx.SetPath("/api/v1/playbook/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(pb.ID)))
	err = s.DeletePlaybook(ctx)
	a.Nil(err)

	pbs, err := models.Playbooks(s.db).All()
	a.Empty(pbs)

	plays, err := models.Jobs(s.db).All()
	a.Empty(plays)

	tasks, err := models.Runs(s.db).All()
	a.Empty(tasks)
	//a.Equal(sql.ErrNoRows, err)
	a.Nil(err)
}

//func TestUpdatePlaybook(t *testing.T) {
//	e := echo.New()
//	a := assert.New(t)
//	s := newTestServer(t)
//
//	failOnError(t, createPlaybook(s, models.Playbook{
//		Name:    "test",
//		GitRepo: "https://github.com/ansible/ansible-examples/tree/master/",
//		Entry:   "tomcat-standalone/site.yml",
//	}))
//
//	plybk, err := models.Playbooks(s.db).One()
//
//	a.Nil(err)
//
//	plybk.Name = "testtest"
//
//	req := newReq(echo.PUT, "/api/v1/playbook", plybk)
//
//	rec := httptest.NewRecorder()
//
//	ctx := e.NewContext(req, rec)
//
//	err = s.UpdatePlaybook(ctx)
//
//	plybks, _ := models.Playbooks(s.db).All()
//
//	if a.Equal(len(plybks), 1) {
//		a.Equal(plybks[0].Name, "testtest")
//	}
//
//	a.Nil(err)
//}
