package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, signinUsecase usecase.SigninUsecase, newsHandler NewsHandler, signinHandler SigninHandler) {
	e.POST("/login", func(c echo.Context) error {
		return config.Login(c, signinUsecase)
	})

	// 以下のルーティングはJWT認証が必要
	r := e.Group("")
	r.Use(middleware.JWTWithConfig(*config.JWTConfig()))

	// news
	r.GET("/news", newsHandler.Get())
	// signin
	r.GET("/signin", signinHandler.Get())
}
