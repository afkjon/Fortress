package handlers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get JWT token from cookie
		cookie, err := c.Cookie("token")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
		}

		// Validate JWT token
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
		}

		// Extract user ID from token claims and set it in the request context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized",
			})
		}

		// Extract the user_id from claims
		userId, ok := claims["user_id"].(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "User ID not found in token.",
			})
		}

		// Store the user_id in the request context
		c.Set("user_id", userId)

		// Call the next handler
		return next(c)
	}
}

func UserStatus(c echo.Context) error {
	userId, parseErr := c.Get("user_id").(string)
	if parseErr {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized",
		})
	}

	// Retrieve user information from the database
	user, err := FindUserById(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "User not found",
		})
	}

	return c.JSON(http.StatusOK, user)
}

func HandleClaims(c echo.Context) error {
	claims := c.Get("token").(*jwt.Token).Claims.(jwt.MapClaims)
	return c.JSON(http.StatusOK, claims)
}
