package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"git.elenet.me/yuelong.huang/pansible/models"

	"crypto/aes"
	"crypto/cipher"
	"github.com/Sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/franela/goreq"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

type Claims struct {
	jwt.StandardClaims
	Task   string `json:"task"`
	Master string `json:"master"`
}

func getUser(ctx echo.Context) *Claims {
	t, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return nil
	}

	c, _ := t.Claims.(*Claims)
	return c
}

func StartServer() {
	app := echo.New()

	s, err := NewServer()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to create a new server")
	}

	go s.startSimpleTLS()

	app.POST("/runnerticket", func(ctx echo.Context) error {
		var body struct {
			ID string `json:"run_uuid"`
		}
		if err := ctx.Bind(&body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest,
				errors.Wrap(err, "Failed to parse request body"))
		}
		if body.ID == "" {
			return echo.NewHTTPError(http.StatusBadRequest,
				errors.New("no run uuid provided"))
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			},
			body.ID,
			fmt.Sprintf("%s:%d", s.Host, s.HTTPPort),
		})

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(s.JWTSecret))
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	})

	s.defineRunnerRoutes(app)

	app.POST("/echo", func(ctx echo.Context) error {
		_, err := io.Copy(os.Stdout, ctx.Request().Body)
		fmt.Println("")
		return err
	})

	s.defineUserRoutes(app)

	logrus.Fatal(app.Start(fmt.Sprintf(":%d", s.HTTPPort)))

}

func (s *Server) startSimpleTLS() {
	// Serves ssh key
	http.HandleFunc("/runner/ssh-key", func(w http.ResponseWriter, req *http.Request) {

		_, err := s.AuthenticateRunner(req)
		if err != nil {
			logrus.WithError(err).Error("failed to authenticate runner")
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		f, err := os.Open(s.SSHKeyFile)
		if err != nil {
			logrus.WithError(err).Error("failed to open ssh keyfile")
			return
		}

		ciphertext, err := ioutil.ReadAll(f)
		if err != nil {
			logrus.WithError(err).Error("failed to read ssh keyfile")
			return
		}

		plaintext, err := s.decryptKey(ciphertext)
		if err != nil {
			logrus.WithError(err).Error("failed to decrypt ssh keyfile")
			return
		}

		_, err = w.Write(plaintext)
		if err != nil {
			logrus.WithError(err).Error("failed to send ssh keyfile content to runner")
			return
		}
	})

	if err := http.ListenAndServeTLS(
		fmt.Sprintf(":%d", s.HTTPSPort),
		"server.crt",
		"server.key",
		nil); err != nil {
		logrus.WithError(err).Fatal("failed to start https server")
	}
}

func (s *Server) defineRunnerRoutes(app *echo.Echo) {

	runner := app.Group("/runner")
	runner.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &Claims{},
		SigningKey: []byte(s.JWTSecret),
	}))

	runner.GET("/inventory/:run_id", s.GetRunInventory)
	runner.GET("/playbook/:job_id", s.GetPlaybookByJobID)
	runner.GET("/run/:uuid", s.FindRunByUUID)

	runner.POST("/events", s.Callback)

}

func (s *Server) defineUserRoutes(app *echo.Echo) {
	v1 := app.Group("/api/v1")
	//v1.Use(s.AuthenticateUser)

	invAPI := v1.Group("/inventory")
	invAPI.POST("", s.CreateInventory)
	invAPI.GET("/:id", s.GetInventory)
	invAPI.GET("/list", s.List(models.TableNames.Inventory))
	invAPI.GET("/search", s.SearchByName(models.TableNames.Inventory))
	invAPI.PUT("/:id", s.UpdateInventory)
	invAPI.DELETE("/:id", s.DeleteInventory)

	playbookAPI := v1.Group("/playbook")
	playbookAPI.POST("", s.CreatePlaybook)
	playbookAPI.GET("/:id", s.GetPlaybook)
	playbookAPI.GET("/list", s.List(models.TableNames.Playbook))
	playbookAPI.GET("/search", s.SearchByName(models.TableNames.Playbook))
	playbookAPI.DELETE("/:id", s.DeletePlaybook)

	jobAPI := v1.Group("/job")
	jobAPI.POST("", s.NewJob)
	jobAPI.GET("/list", s.List(models.TableNames.Job))
	jobAPI.GET("/:uuid", s.GetByUUID(models.TableNames.Job))
	jobAPI.GET("/result/:jobID", s.GetStatsForJob)
	jobAPI.GET("/progress/:jobID", s.GetJobProgress)

	runAPI := v1.Group("/run")
	runAPI.GET("/job/:job_id", s.GetRunsByJobID)
	runAPI.GET("/:uuid", s.GetByUUID(models.TableNames.Run))

	appAPI := v1.Group("/app")
	appAPI.POST("", s.CreateApp)
	appAPI.GET("/list", s.List(models.TableNames.App))
	appAPI.GET("/:appid", s.GetApp)
	appAPI.DELETE("/:appid", s.DeleteApp)

	appAPI.POST("/inventory", s.AddAppInventory)
	appAPI.POST("/playbook", s.AddAppPlaybook)
	appAPI.DELETE("/inventory", s.DeleteAppInventory)
	appAPI.DELETE("/playbook", s.DeleteAppPlaybook)

	//for testing's sake
	app.GET("/api/job/:jobID", s.RerunJob)

}

func placeholder(ctx echo.Context) error {
	return nil
}

func (s *Server) decryptKey(encoded []byte) ([]byte, error) {

	block, err := aes.NewCipher([]byte(s.AESSecret))
	if err != nil {
		return nil, err
	}
	if len(encoded) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := encoded[:aes.BlockSize]
	encoded = encoded[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(encoded, encoded)
	return encoded, nil
}

func (s *Server) AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token, err := ctx.Cookie("SOME_TOKEN")
		if err != nil {
			return ctx.JSON(400, echo.Map{
				"error":   err.Error(),
				"message": "cannot get cookie 'SOME_TOKEN'",
			})
		}

		request := goreq.Request{
			Timeout:     time.Second,
			Method:      http.MethodGet,
			Uri:         "http://some_authentication_service/check_token",
			QueryString: struct{ Token string }{token.Value},
		}

		resp, err := request.Do()
		if err != nil {
			return ctx.JSON(500, echo.Map{
				"error":   err.Error(),
				"message": "failed to send authentication request",
			})
		}

		var body struct {
			User struct {
				WorkCode string `json:"work_code"`
			} `json:"user"`
			Expire int64 `json:"expire_at"`
		}

		if err := resp.Body.FromJsonTo(&body); err != nil {
			return ctx.JSON(401, echo.Map{
				"error":   err.Error(),
				"message": "failed to parse authentication response body",
			})
		}

		if body.User.WorkCode == "" || body.Expire == 0 {
			return ctx.JSON(401, echo.Map{
				"message": "user authentication failed",
			})
		}

		if time.Unix(int64(body.Expire/1000), 0).Before(time.Now()) {
			return ctx.JSON(401, echo.Map{
				"message": "token expired",
			})
		}

		return next(ctx)
	}
}

func (s *Server) AuthenticateRunner(req *http.Request) (*Claims, error) {

	a := req.Header.Get("Authorization")
	l := strings.Split(a, " ")
	if len(l) != 2 {
		return nil, errors.New("failed to get authorization header")
	}
	token := l[1]

	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.JWTSecret), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse jwt token")
	}

	claims, ok := t.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claim content")
	}

	return claims, nil
}
