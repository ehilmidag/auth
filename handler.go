package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "hilmi" && password == "1234" {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Hilmi Dag"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

		refreshToken := jwt.New(jwt.SigningMethodHS256)
		rtClaims := refreshToken.Claims.(jwt.MapClaims)
		rtClaims["sub"] = 1
		rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		rt, err := refreshToken.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"access_token":  t,
			"refresh_token": rt,
		})
	}
	return echo.ErrUnauthorized
}

func (h *Handler) private(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
