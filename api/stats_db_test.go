package api

import (
	"fmt"
	"git.elenet.me/yuelong.huang/pansible/tests"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLatestStatsForJob(t *testing.T) {
	e := echo.New()
	a := assert.New(t)
	s := Server{db: tests.NewTestServer(t).DB}

	tests.InsertPlaybookAndInventory(t)

	req := newReq(echo.GET, "/api/job/result/a127fdc6-de96-4898-88a0-43cad07959d5", nil)

	c := http.Cookie{
		Name:  "COFFEE_TOKEN",
		Value: "9d98762c-4041-4add-8d56-0b75f913cbee",
	}

	req.AddCookie(&c)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	cookie, _ := ctx.Cookie("COFFEE_TOKEN")

	fmt.Println("cookie got from request => ", cookie.Value)

	err := s.GetLatestStatsForJob(ctx)

	fmt.Printf("results => %v", ctx.Response())

	a.Nil(err)
}
