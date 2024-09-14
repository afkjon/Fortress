package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/afkjon/Fortress/hourly/backend/db"
	"github.com/afkjon/Fortress/hourly/backend/models"
	"gorm.io/gorm"
)

var jwtSecret = []byte("your_jwt_secret")

type LoginRequest struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

func Login(c echo.Context) error {
	req := new(LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request."})
	}

	// Set headers
	c.Response().Header().Set("Access-Control-Allow-Origin", os.Getenv("CLIENT_URI"))
	c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

	var user models.User
	if result := db.DB.Where("email = ?", req.Email).First(&user); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}
	if err := user.CheckPassword(req.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}

	token, tokenErr := generateJWT(user)
	if tokenErr != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, map[string]string{"message": "Logged in"})
}

func Register(c echo.Context) error {
	user := new(models.User)

	// Set headers
	c.Response().Header().Set("Access-Control-Allow-Origin", os.Getenv("CLIENT_URI"))
	c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Invalid input."})
	}

	if err := user.HashPassword(user.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	if res := db.DB.Create(&user); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	// Generate JWT
	curUser := models.User{}
	dbErr := db.DB.Where("email = ?", user.Email).First(&curUser)
	if dbErr != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	token, err := generateJWT(curUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	return c.JSON(http.StatusCreated, curUser)
}

func generateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func HandleClaims(c echo.Context) error {
	claims := c.Get("email").(*jwt.Token).Claims.(jwt.MapClaims)
	return c.JSON(http.StatusOK, claims)
}
