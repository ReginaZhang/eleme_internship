package api

import (
	"database/sql"
	"fmt"
	"net/http"
	// import mysql driver
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"

	// Import this so we don't have to use qm.Limit etc.
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"

	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/progress"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Server struct {
	db          *sql.DB
	CurrentJobs progress.Progress
	Host        string `mapstructure:"host"`
	HTTPPort    int    `mapstructure:"http_port"`
	HTTPSPort   int    `mapstructure:"https_port"`
	SSHKeyFile  string `mapstructure:"ssh_key"`
	JWTSecret   string `mapstructure:"jwt_secret"`
	AESSecret   string `mapstructure:"aes_secret"`
}

type DataConfig struct {
	Address  string `mapstructure:"mysql_address"`
	Name     string `mapstructure:"mysql_db"`
	User     string `mapstructure:"mysql_user"`
	Password string `mapstructure:"mysql_pass"`
}

type Config struct {
}

func NewServer() (*Server, error) {
	var s Server
	if err := viper.Unmarshal(&s); err != nil {
		return nil, err
	}

	var dc DataConfig
	if err := viper.Unmarshal(&dc); err != nil {
		return nil, err
	}

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", dc.User, dc.Password, dc.Address, dc.Name),
	)
	if err != nil {
		return nil, err
	}

	s.db = db
	s.CurrentJobs = progress.NewLocal()

	return &s, nil
}

func rc(err error) int {
	if err == sql.ErrNoRows {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

type deployRequest struct {
	AppID         string `json:"appid"`
	Version       string `json:"version"`
	Env           string `json:"env"`
	ConfigVersion string `json:"config_version"`
}

func deployApp(tx boil.Executor, req deployRequest) error {
	app, err := models.Apps(tx, Where("appid = ?", req.AppID)).One()
	if err != nil {
		return echo.NewHTTPError(rc(err),
			errors.Wrap(err, "failed to find app"),
		)
	}

	inv, err := models.Inventories(tx,
		Where("appid = ?", req.AppID),
		Where("env = ?", req.Env),
		Where("version = ?", req.ConfigVersion),
	).One()
	if err != nil {
		return echo.NewHTTPError(rc(err),
			errors.Wrap(err, "failed to find inventory"),
		)
	}

	fmt.Printf("%+v\n", app)
	fmt.Printf("%+v\n", inv)

	job := models.Job{
		Env: req.Env,
		//PlaybookRepo: app.TaskRepo,
		//Playbook:     app.Playbook,
		InventoryID: inv.ID,
	}

	if err := job.Insert(tx); err != nil {
		return errors.Wrap(err, "failed to create play")
	}

	fmt.Printf("%+v\n", job)

	return nil
}

func (s *Server) DeployApp(ctx echo.Context) error {
	var req deployRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			errors.Wrap(err, "faield to parse request body"),
		)
	}

	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin sql tx")
	}

	// config, err := models.Configs(tx, Where(""))

	err = deployApp(tx, req)

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err
}

func rollbackOnErr(err error, tx *sql.Tx) {
	if err != nil && tx != nil {
		if err := tx.Rollback(); err != nil {
			logrus.WithError(err).Errorln("Rollback failed")
		}
	}
}
