package api

import (
	"database/sql"
	"git.elenet.me/yuelong.huang/pansible/models"
	. "git.elenet.me/yuelong.huang/pansible/tx"
	"git.elenet.me/yuelong.huang/pansible/utils"
	"github.com/labstack/echo"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

func (s *Server) CreateApp(ctx echo.Context) error {
	var app models.App

	if err := ctx.Bind(&app); err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "failed to parse app",
		})
	}

	if app.Appid == "" {
		return ctx.JSON(400, echo.Map{
			"error":   "Missing required parameter 'Appid'",
			"message": "failed to validate app",
		})
	}

	if err := app.Insert(s.db); err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "failed to create app",
		})
	}

	return ctx.JSON(201, utils.FilterFields(app))
}

func (s *Server) GetApp(ctx echo.Context) error {
	appid := ctx.Param("appid")
	var t TX
	var err error

	var tx *sql.Tx
	t.Run("start DB transaction", func() error {
		tx, err = s.db.Begin()
		return err
	})

	var app *models.App
	t.Run("get app", func() error {
		app, err = models.Apps(tx, Where("appid = ?", appid)).One()
		return err
	})

	var invNames []string
	t.Run("get inventories of app", func() error {
		invs, err := models.Inventories(tx,
			InnerJoin("app_inventory on app_inventory.inventory_id = inventory.id"),
			Where("app_inventory.app_id = ?", app.ID),
		).All()
		for _, inv := range invs {
			invNames = append(invNames, inv.Name)
		}
		return err
	})

	var pbNames []string
	t.Run("get playbooks of app", func() error {
		pbs, err := models.Playbooks(tx,
			InnerJoin("app_playbook on app_playbook.playbook_id = playbook.id"),
			Where("app_playbook.app_id = ?", app.ID),
		).All()
		for _, pb := range pbs {
			pbNames = append(pbNames, pb.Name)
		}
		return err
	})

	t.Run("commit", tx.Commit)
	rollbackOnErr(t.Error(), tx)

	if t.Error() != nil {
		return ctx.JSON(500, echo.Map{
			"error":   t.Error().Error(),
			"message": "failed to get app",
		})
	}

	res := struct {
		models.App `json:"app"`
		Invs       []string `json:"inventories"`
		Pbs        []string `json:"playbooks"`
	}{*app, invNames, pbNames}

	return ctx.JSON(200, utils.FilterFields(res))
}

func (s *Server) AddAppInventory(ctx echo.Context) error {
	var appInv models.AppInventory
	if err := ctx.Bind(&appInv); err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "failed to parse app_inventory",
		})
	}

	if appInv.AppID == 0 {
		return ctx.JSON(400, echo.Map{
			"error":   "Missing required parameter 'id' of app",
			"message": "failed to validate app_inventory",
		})
	}

	if appInv.InventoryID == 0 {
		return ctx.JSON(400, echo.Map{
			"error":   "Missing required parameter 'inventory_id'",
			"message": "failed to validate app_inventory",
		})
	}

	if err := appInv.Insert(s.db); err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "failed to insert app inventory relationship",
		})
	}

	return ctx.JSON(201, utils.FilterFields(appInv))
}

func (s *Server) AddAppPlaybook(ctx echo.Context) error {
	var appPb models.AppPlaybook
	if err := ctx.Bind(&appPb); err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "failed to parse app_playbook",
		})
	}

	if appPb.AppID == 0 {
		return ctx.JSON(400, echo.Map{
			"error":   "Missing required parameter 'id' of app",
			"message": "failed to validate app_playbook",
		})
	}

	if appPb.PlaybookID == 0 {
		return ctx.JSON(400, echo.Map{
			"error":   "Missing required parameter 'playbook_id'",
			"message": "failed to validate app_playbook",
		})
	}

	if err := appPb.Insert(s.db); err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "failed to insert app playbook relationship",
		})
	}

	return ctx.JSON(201, utils.FilterFields(appPb))
}

func (s *Server) DeleteApp(ctx echo.Context) error {
	appid := ctx.Param("appid")
	if err := models.Apps(s.db, Where("appid = ?", appid)).DeleteAll(); err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "failed to delete app",
		})
	}
	return nil
}

func (s *Server) DeleteAppInventory(ctx echo.Context) error {
	var appInv models.AppInventory
	if err := ctx.Bind(&appInv); err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "failed to parse app_inventory",
		})
	}

	if err := models.AppInventories(s.db,
		Where("app_id = ? AND inventory_id = ?", appInv.AppID, appInv.InventoryID),
	).DeleteAll(); err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "failed to delete app_inventory",
		})
	}
	return nil
}

func (s *Server) DeleteAppPlaybook(ctx echo.Context) error {
	var appPb models.AppPlaybook
	if err := ctx.Bind(&appPb); err != nil {
		return ctx.JSON(400, echo.Map{
			"error":   err.Error(),
			"message": "failed to parse app_playbook",
		})
	}
	if err := appPb.Delete(s.db); err != nil {
		return ctx.JSON(500, echo.Map{
			"error":   err.Error(),
			"message": "failed to delete app_playbook",
		})
	}
	return nil
}
