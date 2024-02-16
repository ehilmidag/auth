package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(
	middleware.JWTConfig{SigningKey: []byte("secret")})

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		user := context.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["admin"].(bool)

		if isAdmin == false {
			return echo.ErrUnauthorized
		}
		return next(context)
	}
}
