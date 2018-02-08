package api

import (
	"encoding/json"
	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/tests"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func createInventory(s *Server, inv models.Inventory) error {
	e := echo.New()
	req := newReq(echo.POST, "/api/v1/inventory", inv)

	ctx := e.NewContext(req, httptest.NewRecorder())

	return s.CreateInventory(ctx)
}

func TestListInventories(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	//	invContent := `[all]
	//adct-devfarm-1.vm.elenet.me
	//`
	//
	//	varBytes, _ := json.Marshal(echo.Map{
	//		"a": 1,
	//	})

	tests.FailOnError(t, createInventory(&s, models.Inventory{
		Env:     "alpha",
		Version: "0.0.1",
		//Hosts:   invContent,
		//Vars:    string(varBytes),
	}))

	inv, err := models.Inventories(s.db).One()
	a.Nil(err)
	a.Equal("#", inv.Hosts)

	req := newReq(echo.GET, "/api/v1/inventory", nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	err = s.GetInventories(ctx)

	a.Equal(http.StatusOK, ctx.Response().Status)
	a.Nil(err)
}

func TestUpdateInventory(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	invContent := `[all]
adct-devfarm-1.vm.elenet.me
`

	varBytes, _ := json.Marshal(echo.Map{
		"a": 1,
	})

	tests.FailOnError(t, createInventory(&s, models.Inventory{Env: "alpha",
		Version: "0.0.1",
		Hosts:   invContent,
		Vars:    string(varBytes)}))

	inv, err := models.Inventories(s.db).One()

	a.Nil(err)

	inv.Env = "beta"

	req := newReq(echo.PUT, "/api/v1/inventory", inv)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	err = s.UpdateInventory(ctx)

	invs, _ := models.Inventories(s.db).All()

	if a.Equal(len(invs), 1) {
		a.Equal(invs[0].Env, "beta")
	}

	a.Nil(err)
}

func TestUpdateInventoryVersion(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	invContent := `[all]
adct-devfarm-1.vm.elenet.me
`

	varBytes, _ := json.Marshal(echo.Map{
		"a": 1,
	})

	tests.FailOnError(t, createInventory(&s, models.Inventory{Env: "alpha",
		Version: "0.0.1",
		Hosts:   invContent,
		Vars:    string(varBytes)}))

	inv, err := models.Inventories(s.db).One()

	a.Nil(err)

	inv.Version = "0.0.2"

	req := newReq(echo.PUT, "/api/v1/inventory", inv)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	err = s.UpdateInventory(ctx)

	invs, _ := models.Inventories(s.db).All()

	if a.Equal(len(invs), 2) {
		a.NotEqual(invs[0].Version, invs[1].Version)
	}

	a.Nil(err)
}

func TestDeleteInventory(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	invContent := `[all]
adct-devfarm-1.vm.elenet.me
`

	varBytes, _ := json.Marshal(echo.Map{
		"a": 1,
	})

	tests.FailOnError(t, createInventory(&s, models.Inventory{Env: "alpha",
		Version: "0.0.1",
		Hosts:   invContent,
		Vars:    string(varBytes)}))

	inv, err := models.Inventories(s.db).One()

	a.Nil(err)

	req := newReq(echo.DELETE, "/", nil)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	ctx.SetPath("/api/v1/inventory/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(inv.ID)))
	err = s.DeleteInventory(ctx)
	a.Nil(err)

	invs, err := models.Inventories(s.db).All()
	a.Empty(invs)
	//a.Equal(sql.ErrNoRows, err)
	a.Nil(err)
}
