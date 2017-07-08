package controller

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/kkimu/blaze-go-app/model"
	"github.com/labstack/echo"
)

const (
	BASE_URL = "http://localhost:8000/static/"
)

type Response struct {
	Here    []model.Video
	Related []model.Video
}

func PostVideo(c echo.Context) error {
	lon := c.FormValue("longitude")
	lat := c.FormValue("latitude")

	here, _, err := getFacility(lon, lat)
	if err != nil {
		return c.JSON(500, err)
	}

	video := model.Video{
		Longitude: lon,
		Latitude:  lat,
		Facility:  here,
	}

	if err := model.InsertVideo(&video); err != nil {
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

func GetVideo(c echo.Context) error {
	lon := c.FormValue("longitude")
	lat := c.FormValue("latitude")

	here, related, err := getFacility(lon, lat)
	if err != nil {
		return c.JSON(500, err)
	}

	hv, err := model.GetVideos(here)
	if err != nil {
		return c.JSON(500, err)
	}
	fmt.Println("hv:", hv)
	rv := []model.Video{}
	for i := range related {
		videos, err := model.GetVideos(related[i])
		if err != nil {
			return c.JSON(500, err)
		}
		for j := range videos {
			rv = append(rv, videos[j])
		}
	}
	res := Response{
		Here:    hv,
		Related: rv,
	}
	pp.Println(res)

	return c.JSON(200, res)
}

func getFacility(longitude string, latitude string) (string, []string, error) {
	related := []string{"disney", "usj"}
	return "colony", related, nil
}
