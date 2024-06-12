package main

import (
	"mariovlv/echo-golang/routes"

	"github.com/labstack/echo/v4"
	prettylogger "github.com/rdbell/echo-pretty-logger"
)

func main() {
	e := echo.New()

	// Serve static files from the "dist" directory
	e.Static("/", "ui/dist")

	// Middlewares
	e.Use(prettylogger.Logger)

	e.GET("api/v1/albums", routes.GetAlbums)
	e.GET("api/v1/albums/:id", routes.GetAlbumByID)

	e.Start(":8080")
}