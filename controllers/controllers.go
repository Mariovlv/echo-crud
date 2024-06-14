package controllers

import (
	"mariovlv/echo-golang/models"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var albums = []models.Album{
	{ID: uuid.New().String(), Title: "Amnesiac", Artist: "Radiohead", Date: time.Date(2001, time.May, 30, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New().String(), Title: "In Rainbows", Artist: "Radiohead", Date: time.Date(2001, time.May, 30, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New().String(), Title: "OK Computer", Artist: "Radiohead", Date: time.Date(2001, time.May, 30, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New().String(), Title: "Kid A", Artist: "Radiohead", Date: time.Date(2001, time.May, 30, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New().String(), Title: "The Bends", Artist: "Radiohead", Date: time.Date(2001, time.May, 30, 0, 0, 0, 0, time.UTC)},
	// Add more albums with random IDs
}

func GetAlbums(c echo.Context) error {
	return c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c echo.Context) error {
	id := c.Param("id")
	for _, album := range albums {
		if album.ID == id {
			return c.JSON(http.StatusOK, album)
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Album not found"})
}

func DeleteUser() {}

func CreateUser() {}
