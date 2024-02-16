package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "hello")
	})
	h := NewHandler()
	e.POST("/login", h.login)
	e.GET("/private", h.private, IsLoggedIn)
	e.GET("/admin", h.private, IsLoggedIn, isAdmin)
	e.Logger.Fatal(e.Start(":1323"))
}
