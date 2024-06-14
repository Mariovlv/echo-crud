package controllers

import (
	"encoding/json"
	"mariovlv/echo-golang/initializers"
	"mariovlv/echo-golang/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAlbums(c echo.Context) error {
	var albums []models.Album

	if err := initializers.DB.Find(&albums).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve albums"})
	}

	return c.JSON(http.StatusOK, albums)
}

func PostAlbum(c echo.Context) error {
	var album models.Album
	json.NewDecoder(c.Request().Body).Decode(&album)

	initializers.DB.Create(&album)

	return c.JSON(http.StatusCreated, album)
}

func GetAlbumByID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid album ID"})
	}

	var album models.Album

	if err := initializers.DB.Where("id = ?", id).First(&album).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Album not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve album"})
	}

	return c.JSON(http.StatusOK, album)
}

func DeleteAlbumById() {}
