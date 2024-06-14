package main

import (
	"mariovlv/echo-golang/initializers"
	"mariovlv/echo-golang/models"
	"mariovlv/echo-golang/routes"

	"github.com/labstack/echo/v4"
	prettylogger "github.com/rdbell/echo-pretty-logger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DBConnection()
	initializers.DB.AutoMigrate(models.Album{})
	initializers.DB.AutoMigrate(models.User{})
}

func main() {
	e := echo.New()
	e.Static("/", "ui/dist")

	e.Use(prettylogger.Logger)

	routes.InitRoutes(e)

	e.Start(":8080")
}
