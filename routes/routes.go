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

	a := e.Group("/auth")
	a.Use(middlewares.RequireAuth)
	a.POST("/login", controllers.Login)
	a.GET("/validate", controllers.Validate)
}
