package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"template/docs"
	"template/internal/db"
	"template/internal/routes"
)

//	@title			Template API
//	@version		1.0
//	@description	Template API

// @BasePath	/
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Env load error: %v\n", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.Example(e)

	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	db.ConnectDatabase()

	e.Logger.Fatal(e.Start(os.Getenv("ECHO_HOST")))
}
