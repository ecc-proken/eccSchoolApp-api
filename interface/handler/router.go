package handler

import (
	"github.com/labstack/echo/v4"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, newsHandler NewsHandler, signinHandler SigninHandler) {

	// news
	e.GET("/news", newsHandler.Get())

	// signin
	e.GET("/signin", signinHandler.Get())
}
