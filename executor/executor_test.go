package executor

import (
	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/tests"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLocalExecutorRun(t *testing.T) {
	a := assert.New(t)
	server := Server{db: tests.NewTestServer(t).DB}
	e := LocalExecutor{}

	run, err := models.Runs(server.db).One()
	a.Nil(err)
	runs := []*models.Run{run}
	err = e.Run(&server, time.Now(), runs)
	a.Nil(err)
}
