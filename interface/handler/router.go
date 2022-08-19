package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

// InitRouting routesの初期化
func InitRouting(
	e *echo.Echo, signinUsecase usecase.SigninUsecase,
	newsHandler NewsHandler,
	signinHandler SigninHandler,
	newsOnlyHandler NewsOnlyHandler,
	calendarHandler CalendarHandler,
) {
	e.POST("/login", func(c echo.Context) error {
		return config.Login(c, signinUsecase)
	})

	// 以下のルーティングはJWT認証が必要
	r := e.Group("")
	r.Use(middleware.JWTWithConfig(*config.JWTConfig()))

	/*
		開発後 e から r に変更する
	*/
	// news
	e.GET("/news", newsHandler.Get())
	// signin
	e.GET("/signin", signinHandler.Get())
	// news-only
	e.GET("/news/:id", newsOnlyHandler.Get())
	// calendar
	e.GET("/calendar/:year/:month", calendarHandler.Get())
}
