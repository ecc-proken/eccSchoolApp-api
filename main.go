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
	// usecase
	signinUsecase := usecase.NewSigninUsecase(signinRepository)
	newsUsecase := usecase.NewNewsUsecase(newsRepository)
	newsOnlyUsecase := usecase.NewNewsOnlyUsecase(newsOnlyRepository)
	// handler
	signinHandler := handler.NewSigninHandler(signinUsecase)
	newsHandler := handler.NewNewsHandler(newsUsecase)
	newsOnlyHandler := handler.NewNewsOnlyHandler(newsOnlyUsecase)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	handler.InitRouting(e, signinUsecase, newsHandler, signinHandler, newsOnlyHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
