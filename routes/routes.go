package routes

import (
	"mariovlv/echo-golang/controllers"
	"mariovlv/echo-golang/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// Public routes
	e.GET("/ping", controllers.Pong) // Ping endpoint for testing

	apiV1 := e.Group("/api/v1")

	// User routes
	apiV1.POST("/signup", controllers.CreateUser)                          // Signup endpoint
	apiV1.POST("/login", controllers.Login)                                // Login endpoint
	apiV1.POST("/validate", controllers.Validate, middlewares.RequireAuth) // Validate token endpoint

	// Album routes
	apiV1.GET("/albums", controllers.GetAlbums, middlewares.RequireAuth)        // Get all albums endpoint
	apiV1.POST("/albums", controllers.PostAlbum)                                // Create album endpoint
	apiV1.GET("/albums/:id", controllers.GetAlbumByID, middlewares.RequireAuth) // Get album by ID endpoint

	// Liked by routes
	apiV1.GET("/likedby/:id", controllers.GetAlbumsLikedBy)                  // Get liked albums endpoint
	apiV1.POST("/likedby", controllers.LikeAlbumBy, middlewares.RequireAuth) // Like album endpoint

	// User profile route
	apiV1.GET("/users/:id", controllers.GetUserByID, middlewares.RequireAuth) // Get user by ID endpoint
}
