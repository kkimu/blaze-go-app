package main

import (
	"fmt"
	"net/http"

	"github.com/kkimu/blaze-go-app/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/videos", postVideo)

	e.Logger.Fatal(e.Start(":8000"))
}

func postVideo(c echo.Context) error {
	fmt.Println(c.Request())
	video := model.Video{
		URL:       "http://test",
		Longitude: c.FormValue("longitude"),
		Latitude:  c.FormValue("latitude"),
	}
	if err := model.InsertVideo(video); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "ok")
}
