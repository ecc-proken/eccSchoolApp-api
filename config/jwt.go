package config

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yumekiti/eccSchoolApp-api/domain"
)

type JwtCustomClaims struct {
	domain.User
	jwt.StandardClaims
}

func MD5(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text + os.Getenv("SALT")))
	return hex.EncodeToString(algorithm.Sum(nil))
}

// {id: string, pw: string}
type Param struct {
	ID       string `json:"id"`
	Password string `json:"pw"`
}

func Login(c echo.Context) error {
	// Bind
	param := new(Param)
	if err := c.Bind(param); err != nil {
		return err
	}

	// Validation
	if param.ID == "" || param.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id or password")
	}
	if len(param.ID) != 7 {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be 7 digits")
	}

	// Authentication

	// Set custom claims
	claims := &JwtCustomClaims{
		User: domain.User {
			ID:       param.ID,
			Password: param.Password,
			UUID:     MD5(param.ID),
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 2147483647,
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
	uuid := claims.User.UUID

	return &domain.User{
		ID:       id,
		Password: password,
		UUID:     uuid,
	}

	// return &domain.User{
	// 	ID:       os.Getenv("TEST_ID"),
	// 	Password: os.Getenv("TEST_PW"),
	// 	UUID:     os.Getenv("TEST_UUID"),
	// }
}
