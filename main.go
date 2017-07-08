package main

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/kkimu/blaze-go-app/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	BASE_URL = "http://localhost:8000/static/"
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

	e.POST("/videos", postVideo)

	e.Logger.Fatal(e.Start(":8000"))
}

func postVideo(c echo.Context) error {
	lon := c.FormValue("longitude")
	lat := c.FormValue("latitude")

	video := model.Video{
		Longitude: lon,
		Latitude:  lat,
	}
	err := model.InsertVideo(&video)
	if err != nil {
		return c.JSON(500, err)
	}
	file, err := c.FormFile("video")
	if err != nil {
		return c.JSON(500, err)
	}
	a := strings.Split(file.Filename, ".")
	fname := strconv.Itoa(video.ID) + "." + a[len(a)-1]
	if err := saveVideo(file, fname); err != nil {
		return c.JSON(500, err)
	}

	video.URL = BASE_URL + fname
	if err := model.UpdateVideo(&video); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, video)
}

func saveVideo(file *multipart.FileHeader, fname string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	err = os.MkdirAll("/static/", 0777)
	if err != nil {
		return err
	}

	// Destination
	dst, err := os.Create("/static/" + fname)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
