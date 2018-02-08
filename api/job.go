package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"git.elenet.me/yuelong.huang/pansible/executor"
	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/progress"
	"git.elenet.me/yuelong.huang/pansible/scheduler"
	. "git.elenet.me/yuelong.huang/pansible/tx"
	"git.elenet.me/yuelong.huang/pansible/utils"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/volatiletech/sqlboiler/queries"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/net/websocket"
	"strings"
)

func (s *Server) NewJob(ctx echo.Context) error {
	var job models.Job
	if err := ctx.Bind(&job); err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "Failed to parse request body",
		})
	}

	if err := ValidateJob(&job); err != nil {
		return ctx.JSON(422, echo.Map{
			"error":   err.Error(),
			"message": "Job validation failed",
		})
	}

	var t TX
	t.Run("insert job", func() error {
		job.UUID = uuid.NewV4().String()
		return job.Insert(s.db)
	})

	var runs []*models.Run
	t.Run("schedule", func() (err error) {
		runs, err = s.ScheduleRuns(&job)
		return err
	})

	t.Run("start job", func() error {
		return s.RunJob(job.UUID, runs)
	})

	if t.Error() != nil {
		job.Status = "failed"
		if err := job.Update(s.db); err != nil {
			return ctx.JSON(500, echo.Map{
				"error":   err.Error(),
				"message": "Failed to update job status",
			})
		}
		s.CurrentJobs.Clean(job.UUID)

		return ctx.JSON(500, echo.Map{
			"error":   t.Error().Error(),
			"message": "Failed to create job",
		})
	}

	return ctx.JSON(201, utils.FilterFields(job))
}

func newScheduler() scheduler.Scheduler {
	return scheduler.NewLocal()
	//return scheduler.NewConstant(5)
}

func (s *Server) ScheduleRuns(job *models.Job) ([]*models.Run, error) {
	sch := newScheduler()

	var t TX

	var tx *sql.Tx
	t.Run("start transaction", func() (err error) {
		tx, err = s.db.Begin()
		return err
	})

	var inv *models.Inventory
	t.Run("get inventory", func() (err error) {
		inv, err = models.FindInventory(tx, job.InventoryID)
		return err
	})

	var runs []*models.Run
	t.Run("schedule", func() (err error) {
		runs, err = sch.Schedule(inv)
		return err
	})

	t.Run("create runs", func() error {
		var rows []string
		for _, run := range runs {
			run.JobID = job.UUID
			run.UUID = uuid.NewV4().String()
			rows = append(rows, fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s')",
				run.UUID,
				run.JobID,
				strings.Replace(run.Env, "'", "\\'", -1),
				strings.Replace(run.Hosts, "'", "\\'", -1),
				strings.Replace(run.Vars, "'", "\\'", -1),
				strings.Replace(run.Limit, "'", "\\'", -1),
			))
		}
		query := fmt.Sprintf("INSERT INTO `%s` (`%s`, `%s`, `%s`, `%s`, `%s`, `%s`) VALUES %s ;",
			models.TableNames.Run,
			models.RunColumns.UUID,
			models.RunColumns.JobID,
			models.RunColumns.Env,
			models.RunColumns.Hosts,
			models.RunColumns.Vars,
			models.RunColumns.Limit,
			strings.Join(rows, ", "),
		)

		_, err := queries.Raw(tx, query).Exec()
		return err
	})

	t.Run("commit", tx.Commit)
	rollbackOnErr(t.Error(), tx)

	return runs, t.Error()
}

func newExecutor(host string) executor.Executor {
	if host == "localhost" {
		return executor.NewLocal()
	}
	return executor.NewAppos()
}

func (s *Server) RunJob(jobID string, runs []*models.Run) error {
	exec := newExecutor(s.Host)
	var err error

	s.CurrentJobs.AddNewJob(jobID, len(runs))
	if err != nil {
		return errors.Wrap(err, "Failed to create a new output for job "+jobID)
	}

	if err := exec.Run(
		fmt.Sprintf("%s:%d", s.Host, s.HTTPPort),
		fmt.Sprintf("%s:%d", s.Host, s.HTTPSPort),
		runs); err != nil {
		return err
	}
	return nil
}

func (s *Server) RerunJob(ctx echo.Context) error {
	jobID := ctx.Param("jobID")
	job, err := models.Jobs(s.db, Where("uuid = ?", jobID)).One()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "Failed to get the job with uuid = "+jobID))
	}
	if job == nil {
		return echo.NewHTTPError(http.StatusNotFound, errors.Wrap(err, "Cannot find job with uuid = "+jobID+" in DB"))
	} else {
		runs, err := models.Runs(s.db, Where("job_id=?", jobID)).All()
		if err != nil {
			return ctx.JSON(500, echo.Map{
				"error":   err.Error(),
				"message": "Failed to get runs of job with job uuid = " + jobID,
			})
		}
		if err := s.RunJob(jobID, runs); err != nil {
			s.CurrentJobs.Clean(job.UUID)
			return ctx.JSON(500, echo.Map{
				"error":   err.Error(),
				"message": "Failed to run the job",
			})
		}
	}
	return ctx.JSON(http.StatusOK, job)
}

func (s *Server) GetJobProgress(ctx echo.Context) error {

	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		jobID := ctx.Param("jobID")
		if err := s.CurrentJobs.Stream(jobID, ws); err != nil {
			logrus.Error("progress streaming failed")
			m := map[string]string{
				"error":   err.Error(),
				"message": "Progress streaming failed; connection closed.",
			}
			if err := progress.SendMessage(ws, m); err != nil {
				logrus.WithError(err).Println("Get job progress failed")
			}
			return
		}

	}).ServeHTTP(ctx.Response(), ctx.Request())

	return nil
}

func decodeJSON(i interface{}, ctx echo.Context) error {
	return json.NewDecoder(ctx.Request().Body).Decode(i)
}

func (s *Server) Callback(ctx echo.Context) error {
	body := make(utils.Map)
	if err := decodeJSON(&body, ctx); err != nil {
		return err
	}

	done, err := s.CurrentJobs.Add(body)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "Failed to process callback",
		})
	}

	if done {
		jobID := body.String("_pansible_job")
		if err := s.saveJobResults(jobID); err != nil {
			return errors.Wrap(err, "failed to save results")
		}
		s.CurrentJobs.Clean(jobID)
	}

	return nil
}

func (s *Server) saveJobResults(jobID string) error {
	var t TX
	var tx *sql.Tx

	summaries, failed, err := s.CurrentJobs.Summaries(jobID)
	if err != nil {
		return errors.Wrap(err, "failed to save job results")
	}

	t.Run("start DB transaction", func() error {
		tx, err = s.db.Begin()
		return err
	})
	for stats, failures := range summaries {
		if stats.Unreachable || (len(failures) != 0) {
			failed = true
		}
		t.Run(fmt.Sprintf("save run stats of %s", stats.Target), func() error {
			return stats.Insert(tx)
		})
		for _, f := range failures {
			f.StatsID = stats.ID
			t.Run(fmt.Sprintf("save failures of %s", stats.Target), func() error {
				return f.Insert(tx)
			})
		}
	}

	t.Run("update job status", func() error {
		status := "successful"
		if failed {
			status = "failed"
		}

		return models.Jobs(s.db, Where("uuid = ?", jobID)).UpdateAll(models.M{"status": status})
	})

	t.Run("commit", tx.Commit)

	rollbackOnErr(t.Error(), tx)
	if t.Error() != nil {
		return errors.Wrap(t.Error(), "failed to save job results and update job status")
	}

	return nil
}
