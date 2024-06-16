package middlewares

import (
	"mariovlv/echo-golang/initializers"
	"mariovlv/echo-golang/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the cookie off request
		cookie, err := c.Cookie("Authorization")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		}
		tokenString := cookie.Value

		// Decode
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Make sure that the token's method conform to what you expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			// Return the secret signing key
			return os.Getenv("SECRET"), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Check the exp and other claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := claims["exp"].(float64)
			if time.Now().Unix() > int64(exp) {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token has expired")
			}

			// Find user with the token sub (subject) claim
			var user models.User
			// You need to implement the findUserByID function to fetch user details from your database
			initializers.DB.First(&user, claims["sub"])
			if user.ID == 0 {
				return echo.NewHTTPError(http.StatusUnauthorized, "User not found")
			}

			// Attach user to context
			c.Set("user", user)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Call the next handler
		return next(c)
	}
}
