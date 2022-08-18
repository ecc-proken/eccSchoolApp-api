package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yumekiti/eccSchoolApp-api/config"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, newsHandler NewsHandler, signinHandler SigninHandler) {

	r := e.Group("")

	r.Use(middleware.JWTWithConfig(*config.JWTConfig()))

	// news
	r.GET("/news", newsHandler.Get())

	// signin
	r.GET("/signin", signinHandler.Get())
}
