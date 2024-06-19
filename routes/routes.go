package routes

import (
	"mariovlv/echo-golang/controllers"
	"mariovlv/echo-golang/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/api/v1/albums", controllers.GetAlbums, middlewares.RequireAuth)
	e.POST("api/v1/albums", controllers.PostAlbum)
	e.GET("api/v1/albums/:id", controllers.GetAlbumByID, middlewares.RequireAuth)
	e.GET("/ping", controllers.Pong)

	e.GET("api/v1/likedby/:id", controllers.GetAlbumsLikedBy)
	e.POST("api/v1/likedby/:user_id/:album_id", controllers.LikeAlbumBy, middlewares.RequireAuth)

	e.GET("api/v1/users/:id", controllers.GetUserByID, middlewares.RequireAuth)

	e.POST("/signup", controllers.CreateUser)
	e.POST("/api/v1/login", controllers.Login)
	e.POST("/validate", controllers.Validate, middlewares.RequireAuth)
}
