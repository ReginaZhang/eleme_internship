package api

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"git.elenet.me/yuelong.huang/pansible/models"
	. "git.elenet.me/yuelong.huang/pansible/tx"
	. "git.elenet.me/yuelong.huang/pansible/utils"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"gopkg.in/yaml.v2"
)

func (s *Server) UploadInvFiles(ctx echo.Context) error {
	env := ctx.FormValue("env")
	version := ctx.FormValue("version")
	name := ctx.FormValue("name")

	hosts, err := getUploadedFile(ctx, "hosts")
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "failed to get hosts file content",
		})
	}
	vars, err := getUploadedFile(ctx, "vars")
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "failed to get vars file content",
		})
	}

	inv := models.Inventory{
		Name:    name,
		Env:     env,
		Version: version,
		Hosts:   hosts,
		Vars:    vars,
	}

	if err := ValidateInventory(&inv); err != nil {
		return ctx.JSON(422, echo.Map{
			"error":   err.Error(),
			"message": "Inventory validation failed",
		})
	}

	if err := inv.Insert(s.db); err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "Failed to insert new inventory",
		})
	}

	return ctx.JSON(http.StatusCreated, inv)
}

func getUploadedFile(ctx echo.Context, name string) (string, error) {
	file, err := ctx.FormFile(name)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("failed to get file %s", name))
	}

	src, err := file.Open()
	if err != nil {
		return "", errors.Wrap(err, "failed to open the file")
	}
	defer src.Close()

	var content []byte

	content, err = ioutil.ReadAll(src)
	if err != nil {
		return "", errors.Wrap(err, "failed to read content of the file")
	}
	return string(content), nil
}

func (s *Server) CreateInventory(ctx echo.Context) error {
	var inv models.Inventory

	if err := ctx.Bind(&inv); err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "Failed to parse request body",
		})
	}

	if err := ValidateInventory(&inv); err != nil {
		return ctx.JSON(422, echo.Map{
			"error":   err.Error(),
			"message": "Inventory validation failed",
		})
	}

	if err := inv.Insert(s.db); err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "Failed to insert new inventory",
		})
	}

	return ctx.JSON(http.StatusCreated, FilterFields(inv))
}

func (s *Server) GetInventory(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "missing or malformed required parameter 'id'",
		})
	}

	inv, err := models.FindInventory(s.db, id)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "Failed to get inventory",
		})
	}

	return ctx.JSON(200, FilterFields(inv))
}

func (s *Server) UpdateInventory(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "missing or malformed required parameter 'id'",
		})
	}

	var inv models.Inventory
	if err := ctx.Bind(&inv); err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "failed to parse request body",
		})
	}

	if err := ValidateInventory(&inv); err != nil {
		return ctx.JSON(422, echo.Map{
			"error":   err.Error(),
			"message": "Inventory validation failed",
		})
	}

	var t TX

	var tx *sql.Tx
	t.Run("start DB transaction", func() error {
		tx, err = s.db.Begin()
		return err
	})

	var old *models.Inventory
	t.Run("find inventory to update", func() error {
		old, err = models.Inventories(tx, Where("id = ?", id)).One()
		return err
	})

	t.Run("udpate inventory", func() error {
		if old.Version != inv.Version {
			inv.ID = 0
			return inv.Insert(tx)
		}
		return inv.Update(tx, "vars", "hosts", "env")
	})

	t.Run("commit", tx.Commit)
	rollbackOnErr(t.Error(), tx)

	if t.Error() != nil {
		return ctx.JSON(500, echo.Map{
			"error":   t.Error().Error(),
			"message": "Failed to update inventory",
		})
	}

	return ctx.JSON(http.StatusOK, FilterFields(inv))
}

func (s *Server) DeleteInventory(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
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

	t.Run("find inventory to delete", func() error {
		return models.Inventories(tx, Where("id = ?", id)).DeleteAll()
	})

	t.Run("commit", tx.Commit)
	rollbackOnErr(t.Error(), tx)

	if t.Error() != nil {
		return ctx.JSON(500, echo.Map{
			"error":   t.Error().Error(),
			"message": "Failed to delete inventory",
		})
	}

	return ctx.JSON(http.StatusNoContent, echo.Map{"inventory_id": id})
}

func (s *Server) GetRunInventory(ctx echo.Context) error {
	runID := ctx.Param("run_id")
	if runID == "" {
		return ctx.JSON(400, echo.Map{
			"error":   "no run id specified",
			"message": "failed to get inventory by run id",
		})
	}

	run, err := models.Runs(s.db, Where("uuid = ?", runID)).One()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": fmt.Sprintf("Failed to get run with runID = %s", runID),
		})
	}

	hosts := make(echo.Map)
	if err := yaml.Unmarshal([]byte(run.Hosts), hosts); err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "Failed to unmarshal hosts",
		})
	}

	vars := make(echo.Map)
	if err := yaml.Unmarshal([]byte(run.Vars), vars); err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "Failed to unmarshal vars",
		})
	}

	return ctx.JSON(http.StatusOK, EnsureStringKey(echo.Map{
		"hosts": hosts,
		"vars":  vars,
	}))
}

func (s *Server) FindInventoryByID(exec boil.Executor, id int64) (*models.Inventory, error) {
	return models.FindInventory(exec, id)
}

func ValidateInventory(inv *models.Inventory) error {

	if inv.Env == "" {
		return errors.New("Missing required parameter 'Env'")
	}

	if inv.Version == "" {
		return errors.New("Missing required parameter 'Version'")
	}

	if err := validateInventoryEnv(inv); err != nil {
		return errors.Wrap(err, "'Env' validation failed")
	}

	if err := validateVars(inv); err != nil {
		return errors.Wrap(err, "'Host' validation failed")
	}

	if err := validateHost(inv); err != nil {
		return errors.Wrap(err, "'Host' validation failed")
	}

	return nil
}

func validateInventoryEnv(inv *models.Inventory) error {
	// TODO
	return nil
}

func validateVars(inv *models.Inventory) error {
	if inv.Vars == "" {
		inv.Vars = "#"
	}
	//TODO
	return nil
}

func validateHost(inv *models.Inventory) error {
	if inv.Hosts == "" {
		inv.Hosts = "#"
	}
	// TODO
	return nil
}
