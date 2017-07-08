package main

import (
	"net/http"

	"github.com/kkimu/blaze-go-app/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "/static")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/videos", controller.PostVideo)
	e.GET("/videos", controller.GetVideo)
	e.Logger.Fatal(e.Start(":8000"))
}
