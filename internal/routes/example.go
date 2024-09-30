package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"template/internal/handlers"
	"template/internal/repository"
)

func Example(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	repositoryHandler := handlers.SetRepositoryHandler(&repository.DefaultRepository{})
	e.GET("/example/:id", repositoryHandler.GetExample)
	e.GET("/example/list", repositoryHandler.ListExamples)
	e.POST("/example", repositoryHandler.CreateExample)
	e.PUT("/example", repositoryHandler.UpdateExample)
	e.DELETE("/example/:id", repositoryHandler.DeleteExample)
}
