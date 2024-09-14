package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/afkjon/Fortress/hourly/backend/db"
	"github.com/afkjon/Fortress/hourly/backend/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func JwtLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Set headers
	c.Response().Header().Set("Access-Control-Allow-Origin", os.Getenv("CLIENT_URI"))
	c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

	var user models.User
	if result := db.DB.Where("email = ?", email).First(&user); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}
	if err := user.CheckPassword(password); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}

	claims := &jwtCustomClaims{
		Email: email,
		Admin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecret)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}
