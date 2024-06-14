package routes

import (
	"mariovlv/echo-golang/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/api/v1/albums", controllers.GetAlbums)
	e.GET("api/v1/albums/:id", controllers.GetAlbumByID)
}