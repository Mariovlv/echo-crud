package controllers

import (
	"encoding/json"
	"fmt"
	"mariovlv/echo-golang/helpers"
	"mariovlv/echo-golang/initializers"
	"mariovlv/echo-golang/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUserByID(c echo.Context) error {
	userID := c.Param("id")

	var user models.User
	if err := initializers.DB.Preload("Albums").First(&user, "id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	return c.JSON(http.StatusOK, user)
}

func GetAlbums(c echo.Context) error {
	//helpers.WriteCookie(c)
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

func Pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

type ErrorResponse struct {
	Err string
}

func CreateUser(c echo.Context) error {
	// Get the email, username, password from req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		return c.String(http.StatusBadRequest, "Error to read the body")
	}
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return c.String(http.StatusBadRequest, "failing hashing the password")
	}
	// create the user, write on the db
	user := models.User{Password: string(hash), Email: body.Email}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, "error: user already exists with same  email")
	}

	//Respond
	return c.String(http.StatusCreated, "User created correctly")
}

func Login(c echo.Context) error {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		return c.String(http.StatusBadRequest, "Error to read the email and password")
	}

	// loog up request user in DB
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		return c.String(http.StatusBadRequest, "error: incorrect user")
	}

	// compare sent password with the already hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		return c.String(http.StatusBadRequest, "error: incorrect user")
	}
	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return c.String(http.StatusBadRequest, "error: failed generating token")
	}

	fmt.Println(tokenString)

	helpers.WriteCookie(c, tokenString)

	return c.String(http.StatusCreated, "User authenticated correctly")
}

func Validate(c echo.Context) error {
	fmt.Println(c)

	return c.String(http.StatusOK, "Im logged in")
}

// GET
func GetAlbumsLikedBy(c echo.Context) error {
	var album models.Album
	albumID := c.Param("id")

	if err := initializers.DB.Preload("Users").First(&album, "id = ?", albumID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Album not found"})
	}

	return c.JSON(http.StatusOK, album.Users)
}

type LikeAlbumRequest struct {
	UserID  string `json:"user_id"`
	AlbumID string `json:"album_id"`
}

func LikeAlbumBy(c echo.Context) error {
	var req LikeAlbumRequest

	// Parse the request body
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	// Validate userID and albumID
	if req.UserID == "" || req.AlbumID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID or album ID"})
	}

	user, err := findUserByID(req.UserID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	album, err := findAlbumByID(req.AlbumID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Album not found"})
	}

	if err := likeAlbum(&user, &album); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not like album"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Album liked successfully"})
}

func findUserByID(userID string) (models.User, error) {
	var user models.User
	if err := initializers.DB.First(&user, "id = ?", userID).Error; err != nil {
		return user, err
	}
	return user, nil
}

func findAlbumByID(albumID string) (models.Album, error) {
	var album models.Album
	if err := initializers.DB.First(&album, "id = ?", albumID).Error; err != nil {
		return album, err
	}
	return album, nil
}

func likeAlbum(user *models.User, album *models.Album) error {
	// Start a transaction
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(user).Association("LikedAlbums").Append(album); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
