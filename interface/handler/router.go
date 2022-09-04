package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yumekiti/eccSchoolApp-api/config"
)

// InitRouting routesの初期化
func InitRouting(
	e *echo.Echo,
	newsHandler NewsHandler,
	signinHandler SigninHandler,
	newsOnlyHandler NewsOnlyHandler,
	calendarHandler CalendarHandler,
	attendanceHandler AttendanceHandler,
	timetableHandler TimetableHandler,
) {
	e.POST("/signin", func(c echo.Context) error {
		return config.Login(c)
	})

	// 以下のルーティングはJWT認証が必要
	r := e.Group("")
	r.Use(middleware.JWTWithConfig(*config.JWTConfig()))

	r.GET("/uuid", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"uuid": config.GetUser(c).UUID})
	})
	// news
	r.GET(":uuid/news", newsHandler.Get())
	// signin
	r.GET(":uuid/signin", signinHandler.Get())
	// news-only
	r.GET(":uuid/news/:id", newsOnlyHandler.Get())
	// calendar
	r.GET(":uuid/calendar/:year/:month", calendarHandler.Get())
	// attendance
	r.GET(":uuid/attendance", attendanceHandler.Get())
	// timetable
	r.GET(":uuid/timetable/:week", timetableHandler.Get())
}
