package api

import (
	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/tests"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func createJob(s *Server, job models.Job) error {
	e := echo.New()
	req := newReq(echo.POST, "/api/v1/job", job)

	ctx := e.NewContext(req, httptest.NewRecorder())
	var body struct {
		models.Job
		DoNotRun bool `json:"do_not_run"`
	}
	if err := ctx.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			errors.Wrap(err, "Failed to parse request body"),
		)
	}

	if err := ValidateJob(&body.Job); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity,
			errors.Wrap(err, "Job validation failed"))
	}

	body.UUID = uuid.NewV4().String()

	if err := body.Insert(s.db); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			errors.Wrap(err, "Failed to insert new Job"))
	}
	return nil
}

func TestListJobs(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	tests.FailOnError(t, createJob(&s, models.Job{
		InventoryID: 1,
		PlaybookID:  1,
		Env:         "for testing's sake",
	}))

	req := newReq(echo.GET, "/api/v1/job", nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	err := s.ListJobs(ctx)

	a.Equal(http.StatusOK, ctx.Response().Status)
	a.Nil(err)
}

func TestDeleteJob(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	tests.FailOnError(t, createJob(&s, models.Job{
		InventoryID: 1,
		PlaybookID:  1,
		Env:         "for testing's sake",
	}))

	tests.FailOnError(t, createRun(&s, models.Run{
		JobID: "1",
		Env:   "for testing's sake",
	}))

	job, err := models.Jobs(s.db).One()

	a.Nil(err)

	req := newReq(echo.DELETE, "/", nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	ctx.SetPath("/api/v1/job/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(job.ID)))
	err = s.DeleteJob(ctx)
	a.Nil(err)

	jobs, err := models.Jobs(s.db).All()
	a.Empty(jobs)

	tasks, err := models.Runs(s.db).All()
	a.Empty(tasks)
	//a.Equal(sql.ErrNoRows, err)
	a.Nil(err)
}

//func TestUpdateJob(t *testing.T) {
//	e := echo.New()
//	a := assert.New(t)
//	s := newTestServer(t)
//
//	failOnError(t, createJob(s, models.Job{
//		InventoryID:    1,
//		PlaybookID: 1,
//		Env:   "for testing's sake",
//	}))
//
//	job, err := models.Jobs(s.db).One()
//
//	a.Nil(err)
//
//	job.PlaybookID = 2
//
//	req := newReq(echo.PUT, "/api/v1/job", job)
//
//	rec := httptest.NewRecorder()
//
//	ctx := e.NewContext(req, rec)
//
//	err = s.UpdateJob(ctx)
//
//	jobs, _ := models.Jobs(s.db).All()
//
//	if a.Equal(len(jobs), 1) {
//		a.Equal(jobs[0].JobbookID, 2)
//	}
//
//	a.Nil(err)
//}
