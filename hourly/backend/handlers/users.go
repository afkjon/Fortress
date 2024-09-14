package handlers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/afkjon/Fortress/hourly/backend/db"
	"github.com/afkjon/Fortress/hourly/backend/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func FindUserById(id string) (models.User, error) {
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func Login(c echo.Context) error {
	req := new(LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request."})
	}

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
		c.Logger().Error(tokenErr)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Logged in",
	})
}

func Register(c echo.Context) error {
	user := new(models.User)

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
		c.Logger().Error(err)
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
	claims := jwt.MapClaims{
		"user_id": strconv.FormatUint(uint64(user.ID), 10),
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
