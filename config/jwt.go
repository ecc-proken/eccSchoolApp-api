package config

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

type JwtCustomClaims struct {
	domain.User
	jwt.StandardClaims
}

func Login(c echo.Context, signinUsecase usecase.SigninUsecase) error {
	// Bind
	id := c.FormValue("id")
	password := c.FormValue("password")

	// check id and password
	getSignin, err := signinUsecase.Get(&domain.User{
		Id:       id,
		Password: password,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if getSignin.Status != 200 {
		return c.JSON(http.StatusBadRequest, getSignin.Message)
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		User: domain.User{
			Id:       id,
			Password: password,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func JWTConfig() *middleware.JWTConfig {
	return &middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
}

func GetUser(c echo.Context) *domain.User {
	// tokenからユーザー情報を取得
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(*JwtCustomClaims)
	// id := claims.User.Id
	// password := claims.User.Password

	// return &domain.User{
	// 	Id:       id,
	// 	Password: password,
	// }

	// 開発後上のコメントアウトを有効にし下記は削除
	return &domain.User{
		Id:       os.Getenv("TEST_ID"),
		Password: os.Getenv("TEST_PW"),
	}
}
