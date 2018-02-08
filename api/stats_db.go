package api

import (
	"database/sql"
	"net/http"

	"git.elenet.me/yuelong.huang/pansible/models"
	. "git.elenet.me/yuelong.huang/pansible/tx"
	. "git.elenet.me/yuelong.huang/pansible/utils"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (s *Server) GetStatsForJob(ctx echo.Context) error {

	jobID := ctx.Param("jobID")
	var t TX
	var tx *sql.Tx
	var err error
	var allStats models.StatisticSlice
	res := make(map[string]map[string]interface{})

	t.Run("start DB transaction", func() error {
		tx, err = s.db.Begin()
		return err
	})

	t.Run("get statistics", func() error {
		allStats, err = models.Statistics(tx,
			qm.InnerJoin("run on statistics.run_id = run.uuid"),
			qm.Where("run.job_id = ?", jobID)).All()
		return err
	})

	for _, st := range allStats {
		t.Run("get failures", func() error {
			failures, err := models.Failures(tx, qm.Where("stats_id = ?", st.ID)).All()
			if err != nil {
				return err
			}

			res[st.Target] = map[string]interface{}{
				"stats":    st,
				"failures": failures,
			}

			return nil
		})
	}

	t.Run("commit", tx.Commit)

	rollbackOnErr(t.Error(), tx)
	if t.Error() != nil {
		return ctx.JSON(500, echo.Map{
			"error":   t.Error().Error(),
			"message": "Failed to get stats for job",
		})
	}

	return ctx.JSON(http.StatusOK, FilterFields(res))
}
