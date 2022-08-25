package main

import (
	"fmt"

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
	// usecase
	signinUsecase := usecase.NewSigninUsecase(signinRepository)
	newsUsecase := usecase.NewNewsUsecase(newsRepository)
	newsOnlyUsecase := usecase.NewNewsOnlyUsecase(newsOnlyRepository)
	calendarUsecase := usecase.NewCalendarUsecase(calendarRepository)
	attendanceUsecase := usecase.NewAttendanceUsecase(attendanceRepository)
	// handler
	signinHandler := handler.NewSigninHandler(signinUsecase)
	newsHandler := handler.NewNewsHandler(newsUsecase)
	newsOnlyHandler := handler.NewNewsOnlyHandler(newsOnlyUsecase)
	calendarHandler := handler.NewCalendarHandler(calendarUsecase)
	attendanceHandler := handler.NewAttendanceHandler(attendanceUsecase)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	handler.InitRouting(e,
		newsHandler,
		signinHandler,
		newsOnlyHandler,
		calendarHandler,
		attendanceHandler,
	)

	fmt.Println("Server is running on port :" + config.GetEnv("PORT", "8080"))

	// Start server
	e.Logger.Fatal(e.Start(":" + config.GetEnv("PORT", "8080")))
}
