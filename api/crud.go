package api

import (
	"fmt"
	"strconv"

	"git.elenet.me/yuelong.huang/pansible/models"
	. "git.elenet.me/yuelong.huang/pansible/utils"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func all(tableName string, exe boil.Executor, qm ...qm.QueryMod) (interface{}, error) {
	switch tableName {
	case models.TableNames.Inventory:
		return models.Inventories(exe, qm...).All()
	case models.TableNames.Playbook:
		return models.Playbooks(exe, qm...).All()
	case models.TableNames.Job:
		return models.Jobs(exe, qm...).All()
	default:
		return nil, errors.Errorf("table not found")
	}
}

func (s *Server) List(tableName string) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		offset, _ := strconv.Atoi(ctx.QueryParam("offset"))
		limit, _ := strconv.Atoi(ctx.QueryParam("num"))
		if limit == 0 || limit > 100 {
			limit = 100
		}

		res, err := all(tableName, s.db, qm.Offset(offset), qm.Limit(limit))

		if err != nil {
			return ctx.JSON(500, echo.Map{
				"error":   err.Error(),
				"message": fmt.Sprintf("Failed to list %s", tableName),
			})
		}

		return ctx.JSON(200, FilterFields(res))
	}
}

func (s *Server) SearchByName(tableName string) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		if name == "" {
			return ctx.JSON(400, echo.Map{
				"message": "Missing name",
			})
		}

		res, err := all(tableName, s.db,
			qm.Where("name like ?", fmt.Sprintf("%s%%", name)),
			qm.Limit(100),
		)
		if err != nil {
			return ctx.JSON(500, echo.Map{
				"error":   err.Error(),
				"message": fmt.Sprint("Failed to search %s", tableName),
			})
		}

		return ctx.JSON(200, FilterFields(res))
	}
}

func (s *Server) GetByID(tableName string) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		idS := ctx.Param("id")
		id, err := strconv.Atoi(idS)
		if err != nil || id == 0 {
			return ctx.JSON(400, echo.Map{
				"message": "Missing id",
			})
		}

		res, err := all(tableName, s.db,
			qm.Where("id = ?", id),
		)
		if err != nil {
			return ctx.JSON(500, echo.Map{
				"error":   err.Error(),
				"message": fmt.Sprint("Failed to get %s by ID", tableName),
			})
		}
		return ctx.JSON(200, FilterFields(res))
	}
}

func (s *Server) GetByUUID(tableName string) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		uuid := ctx.Param("uuid")
		if uuid == "" {
			return ctx.JSON(400, echo.Map{
				"message": "Missing uuid",
			})
		}

		res, err := all(tableName, s.db,
			qm.Where("uuid = ?", uuid),
		)
		if err != nil {
			return ctx.JSON(500, echo.Map{
				"error":   err.Error(),
				"message": fmt.Sprint("Failed to get %s by UUID", tableName),
			})
		}
		return ctx.JSON(200, FilterFields(res))
	}
}
