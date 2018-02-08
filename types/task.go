package types

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Claims struct {
	jwt.StandardClaims
	Master string `json:"master"`
	Task   string `json:"task"`
	//Env    string `json:"env"`
}

//func ParseClaims(token string) (*Claims, error) {
//	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		secret, _ := viper.GetStringMap("server")["jwt_secret"].(string)
//		fmt.Printf("\nin parsing claims secret => %s\n\n", secret)
//		return []byte(secret), nil
//	})
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to parse jwt token")
//	}
//
//	claims, ok := t.Claims.(*Claims)
//	if !ok {
//		return nil, errors.New("invalid claim content")
//	}
//
//	return claims, nil
//}

type Task struct {
	TaskID string
	Type   string
	URL    string
	Entry  string
}

type Config struct {
	AppID string
	Body  []byte
}

type App struct {
	AppID   string
	GitRepo string
	Version string
}
