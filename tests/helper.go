package tests

import (
	"database/sql"
	"fmt"
	"git.elenet.me/yuelong.huang/pansible/models"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

type Server struct {
	DB *sql.DB
}

func InsertPlaybookAndInventory(t *testing.T) {
	s := NewTestServer(t)
	pb := models.Playbook{
		Name:    "test",
		GitRepo: "git@some_repo",
		Entry:   "some_playbook.yml",
	}
	err := pb.Insert(s.DB)
	if err != nil {
		fmt.Println("error inserting playbook => ", err.Error())
	}

	content, err := ioutil.ReadFile("../tmp/inv.yml")
	if err != nil {
		fmt.Println("error reading host file => ", err.Error())
	}

	fmt.Println("host string => ", string(content))
	inv := models.Inventory{
		Env:     "alpha",
		Version: "0.0.1",
		Hosts:   string(content),
		Vars:    "---",
	}
	err = inv.Insert(s.DB)
	if err != nil {
		fmt.Println("error inserting inventory => ", err.Error())
	}
}

func runSqlScript(filename string, options ...string) error {
	options = append(options, "-h", "127.0.0.1", "-uroot", "-ptoor")
	cmd := exec.Command("mysql", options...)
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	cmd.Stdin = f
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
	}

	return err
}

func flushDB() error {
	if err := runSqlScript("../drop.sql"); err != nil {
		return err
	}

	if err := runSqlScript("../create.sql", "-Dpansible"); err != nil {
		return err
	}

	return nil
}

func FailOnError(t *testing.T, err error) {
	if err == nil {
		return
	}

	t.Error(err)
	t.FailNow()
}

func NewTestServer(t *testing.T) *Server {
	db, err := sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/pansible?charset=utf8&parseTime=True")
	FailOnError(t, err)
	FailOnError(t, flushDB())

	return &Server{DB: db}
}
