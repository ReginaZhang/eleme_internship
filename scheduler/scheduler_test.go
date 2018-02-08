package scheduler

import (
	"git.elenet.me/yuelong.huang/pansible/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestLocalScheduler_Schedule(t *testing.T) {
	a := assert.New(t)
	s := LocalScheduler{}
	content, err := ioutil.ReadFile("/tmp/test.yml")
	a.Nil(err)
	inv := models.Inventory{
		Hosts: string(content),
	}
	s.Schedule(&inv)

}
