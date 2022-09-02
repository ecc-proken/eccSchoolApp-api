package config

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yumekiti/eccSchoolApp-api/domain"
)

type JwtCustomClaims struct {
	domain.User
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	// Bind
	id := c.FormValue("id")
	password := c.FormValue("password")

	// Set custom claims
	claims := &JwtCustomClaims{
		User: domain.User{
			ID:       id,
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	id := claims.User.ID
	password := claims.User.Password

	return &domain.User{
		ID:       id,
		Password: password,
	}

	// return &domain.User{
	// 	Id:       os.Getenv("TEST_ID"),
	// 	Password: os.Getenv("TEST_PW"),
	// }
}
