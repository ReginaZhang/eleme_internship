package client

import (
	"fmt"
	"git.elenet.me/yuelong.huang/pansible/api"
	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/tests"
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetTask(t *testing.T) {
	s := tests.NewTestServer(t)
	a := assert.New(t)

	runID := uuid.NewV4().String()
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, api.Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
		runID,
		"localhost:5757",
		"test",
	})

	token, err := claim.SignedString([]byte("secret"))
	a.Nil(err)

	client, err := New(token)
	a.Nil(err)

	task := models.Run{
		JobID: 1,
		Env:   "for testing's sake",
		UUID:  runID,
	}

	err = task.Insert(s.DB)
	a.Nil(err)

	newRun, err := client.GetRun(runID)

	a.Nil(err)
	a.Equal(runID, newRun.UUID)

	fmt.Println(newRun)
}
