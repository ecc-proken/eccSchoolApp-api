package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/infrastructure"
	"github.com/yumekiti/eccSchoolApp-api/interface/handler"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

func init() {
	config.LoadEnv()
}

func main() {
	// repository
	signinRepository := infrastructure.NewSigninRepository()
	newsRepository := infrastructure.NewNewsRepository()
	newsOnlyRepository := infrastructure.NewNewsOnlyRepository()
	calendarRepository := infrastructure.NewCalendarRepository()
	attendanceRepository := infrastructure.NewAttendanceRepository()
	timetableRepository := infrastructure.NewTimetableRepository()
	// usecase
	signinUsecase := usecase.NewSigninUsecase(signinRepository)
	newsUsecase := usecase.NewNewsUsecase(newsRepository)
	newsOnlyUsecase := usecase.NewNewsOnlyUsecase(newsOnlyRepository)
	calendarUsecase := usecase.NewCalendarUsecase(calendarRepository)
	attendanceUsecase := usecase.NewAttendanceUsecase(attendanceRepository)
	timetableUsecase := usecase.NewTimetableUsecase(timetableRepository)
	// handler
	signinHandler := handler.NewSigninHandler(signinUsecase)
	newsHandler := handler.NewNewsHandler(newsUsecase)
	newsOnlyHandler := handler.NewNewsOnlyHandler(newsOnlyUsecase)
	calendarHandler := handler.NewCalendarHandler(calendarUsecase)
	attendanceHandler := handler.NewAttendanceHandler(attendanceUsecase)
	timetableHandler := handler.NewTimetableHandler(timetableUsecase)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
	}))

	// Routes
	handler.InitRouting(e,
		newsHandler,
		signinHandler,
		newsOnlyHandler,
		calendarHandler,
		attendanceHandler,
		timetableHandler,
	)

	// Start server
	e.Logger.Fatal(e.Start(":" + config.GetEnv("PORT", "8080")))
}
