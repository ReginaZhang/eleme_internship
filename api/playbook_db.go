package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"git.elenet.me/yuelong.huang/pansible/models"
	. "git.elenet.me/yuelong.huang/pansible/tx"
	. "git.elenet.me/yuelong.huang/pansible/utils"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

func (s *Server) CreatePlaybook(ctx echo.Context) error {
	var pb models.Playbook

	if err := ctx.Bind(&pb); err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "Failed to parse request body",
		})
	}

	if err := ValidatePlaybook(&pb); err != nil {
		return ctx.JSON(422, echo.Map{
			"error":   err.Error(),
			"message": "Playbook validation failed",
		})
	}

	if err := pb.Insert(s.db); err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "Failed to insert new Playbook",
		})
	}

	ctx.Response().Status = http.StatusCreated

	return nil
}

func (s *Server) GetPlaybook(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "invalid id",
		})
	}

	var t TX

	var tx *sql.Tx
	t.Run("start transaction", func() error {
		tx, err = s.db.Begin()
		return err
	})

	var res echo.Map
	t.Run("find playbook", func() error {
		pb, err := models.FindPlaybook(tx, id)
		if err == nil {
			res, _ = FilterFields(pb).(map[string]interface{})
		}
		return err
	})

	t.Run("find inventory", func() error {
		invs, err := models.Inventories(tx,
			InnerJoin("playbook_inventory pbinv on pbinv.playbook = ?", id),
			Where("inventory.id = pbinv.inventory"),
		).All()

		if err == nil {
			res["inventories"] = FilterFields(invs, "vars", "hosts")

		}
		return err
	})

	t.Run("commit", tx.Commit)
	rollbackOnErr(t.Error(), tx)

	if t.Error() != nil {
		return ctx.JSON(500, echo.Map{
			"error":   t.Error().Error(),
			"message": "Failed to get playbook",
		})
	}

	return ctx.JSON(200, res)
}

func (s *Server) DeletePlaybook(ctx echo.Context) error {
	id := ctx.Param("id")
	numid, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "Failed to parse url param 'id' ",
		})
	}

	var t TX
	var tx *sql.Tx

	t.Run("start DB transaction", func() error {
		tx, err = s.db.Begin()
		return err
	})

	t.Run("find playbook to delete", func() error {
		return models.Playbooks(tx, Where("id = ?", numid)).DeleteAll()
	})

	t.Run("commit", tx.Commit)

	rollbackOnErr(t.Error(), tx)
	if t.Error() != nil {
		return ctx.JSON(500, echo.Map{
			"error":   t.Error().Error(),
			"message": "Failed to delete playbook",
		})
	}

	return ctx.JSON(http.StatusNoContent, map[string]int{"playbook_id": numid})
}

func (s *Server) GetPlaybookByJobID(ctx echo.Context) error {
	jobID := ctx.Param("job_id")
	if jobID != "" {
		playbook, err := models.Playbooks(s.db, InnerJoin("job on job.playbook_id = playbook.id"), Where("job.uuid = ?", jobID)).One()
		if err != nil {
			return ctx.JSON(500, echo.Map{
				"error":   err.Error(),
				"message": fmt.Sprintf("Failed to find playbook by jobID with jobID = %s", jobID),
			})
		}
		return ctx.JSON(http.StatusOK, playbook)
	}
	return nil
}

func (s *Server) FindPlaybookByID(exec boil.Executor, id int64) (*models.Playbook, error) {
	pb, err := models.FindPlaybook(exec, id)
	return pb, err
}

func ValidatePlaybook(pb *models.Playbook) error {

	if pb.Name == "" {
		return errors.New("Missing required parameter 'Appid'")
	}

	if pb.GitRepo == "" {
		return errors.New("Missing required parameter 'Env'")
	}

	if pb.Entry == "" {
		return errors.New("Missing required parameter 'Version'")
	}

	return nil
}
