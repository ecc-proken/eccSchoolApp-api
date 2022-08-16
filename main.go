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
	newsRepository := infrastructure.NewNewsRepository()
	newsUsecase := usecase.NewNewsUsecase(newsRepository)
	newsHandler := handler.NewNewsHandler(newsUsecase)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.InitRouting(e, newsHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
