package routes

import (
	"mariovlv/echo-golang/controllers"
	"mariovlv/echo-golang/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/api/v1/albums", controllers.GetAlbums)
	e.POST("api/v1/albums", controllers.PostAlbum)
	e.GET("api/v1/albums/:id", controllers.GetAlbumByID)
	e.GET("/ping", controllers.Pong)

	e.POST("/signup", controllers.CreateUser)
	e.POST("/login", controllers.Login)
	e.POST("/validate", controllers.Validate, middlewares.RequireAuth)
}
