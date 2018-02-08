package api

import (
	"fmt"

	"git.elenet.me/yuelong.huang/pansible/models"
	. "git.elenet.me/yuelong.huang/pansible/utils"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

func (s *Server) GetRunsByJobID(ctx echo.Context) error {
	jobID := ctx.Param("job_id")
	if jobID == "" {
		return ctx.JSON(400, echo.Map{
			"error":   "no job_id provided",
			"message": "failed to get runs by job id",
		})
	}

	runs, err := models.Runs(s.db, Where("job_id = ?", jobID)).All()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "failed to get runs from DB",
		})
	}

	return ctx.JSON(200, FilterFields(runs))

}

func (s *Server) FindRunByUUID(ctx echo.Context) error {
	id := ctx.Param("uuid")
	if id == "" {
		return ctx.JSON(400, echo.Map{
			"error":   "no uuid specified",
			"message": "failed to get run by uuid",
		})
	}
	i, err := models.Runs(s.db, Where("uuid = ?", id)).One()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": fmt.Sprintf("failed to find the Run with uuid = %s", id),
		})
	}
	return ctx.JSON(200, FilterFields(i))
}

func (s *Server) FindRunByID(exec boil.Executor, id int64) (*models.Run, error) {
	run, err := models.FindRun(exec, id)
	return run, err
}

func ValidateRun(run *models.Run) error {

	if run.JobID == "" {
		return errors.New("Missing required parameter 'JobID'")
	}

	if run.Env == "" {
		return errors.New("Missing required parameter 'Env'")
	}

	if err := validateRunEnv(run); err != nil {
		return errors.Wrap(err, "'Env' validation failed")
	}

	return nil
}

func validateRunEnv(run *models.Run) error {
	// TODO
	return nil
}
