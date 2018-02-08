package api

import (
	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/tests"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func createRun(s *Server, run models.Run) error {
	e := echo.New()
	req := newReq(echo.POST, "/api/v1/run", run)

	ctx := e.NewContext(req, httptest.NewRecorder())

	if err := ctx.Bind(&run); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			errors.Wrap(err, "Failed to parse request body"),
		)
	}

	if err := ValidateRun(&run); err != nil {
		return err
	}

	if err := run.Insert(s.db); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			errors.Wrap(err, "Failed to insert new Run"))
	}

	ctx.Response().Status = http.StatusCreated

	return nil

}

func TestListRuns(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	tests.FailOnError(t, createRun(&s, models.Run{
		JobID: "6290098d-c949-49b0-813e-2ce03d4a6a1d",
		Env:   "for testing's sake",
	}))

	req := newReq(echo.GET, "/api/v1/run", nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	err := s.ListRuns(ctx)

	a.Equal(http.StatusOK, ctx.Response().Status)
	a.Nil(err)
}

func TestDeleteRun(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	tests.FailOnError(t, createRun(&s, models.Run{
		JobID: "6290098d-c949-49b0-813e-2ce03d4a6a1d",
		Env:   "for testing's sake",
	}))

	run, err := models.Runs(s.db).One()

	a.Nil(err)

	req := newReq(echo.DELETE, "/", nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	ctx.SetPath("/api/v1/run/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(run.ID)))
	err = s.DeleteRun(ctx)
	a.Nil(err)

	runs, err := models.Runs(s.db).All()
	a.Empty(runs)
	//a.Equal(sql.ErrNoRows, err)
	a.Nil(err)
}

func TestUpdateRun(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	tests.FailOnError(t, createRun(&s, models.Run{
		JobID: "6290098d-c949-49b0-813e-2ce03d4a6a1d",
		Env:   "for testing's sake",
	}))

	run, err := models.Runs(s.db).One()

	a.Nil(err)

	run.JobID = "2"

	req := newReq(echo.PUT, "/api/v1/run", run)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	err = s.UpdateRun(ctx)

	runs, _ := models.Runs(s.db).All()

	if a.Equal(1, len(runs)) {
		a.Equal("2", runs[0].JobID)
	}

	a.Nil(err)
}
