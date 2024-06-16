package helpers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func WriteCookie(c echo.Context, cook string) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "Authorization", Value: cook, Expires: expiration, HttpOnly: true, Secure: false}
	c.SetCookie(&cookie)

}
