package controller

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"github.com/kkimu/blaze-go-app/model"
	"github.com/labstack/echo"
	"github.com/opennota/screengen"
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
	tfname := strconv.Itoa(video.ID) + "_thumbnail.jpg"
	if err := saveVideo(file, fname); err != nil {
		return c.JSON(500, err)
	}
	go func() {
		img, err := generateThumbnail("/static/" + fname)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := saveImage(img, tfname); err != nil {
			fmt.Println(err)
			return
		}
	}()

	video.URL = BASE_URL + fname
	video.ThumbnailURL = BASE_URL + tfname
	if err := model.UpdateVideo(&video); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, video)
}

func generateThumbnail(fn string) (image.Image, error) {
	generator, err := screengen.NewGenerator(fn)
	if err != nil {
		return nil, err
	}
	img, err := generator.Image(1000)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func saveImage(img image.Image, fname string) error {
	dst, err := os.Create("/static/" + fname)
	if err != nil {
		return err
	}

	option := &jpeg.Options{Quality: 100}

	if err = jpeg.Encode(dst, img, option); err != nil {
		return err
	}
	return nil
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
	rv := []model.Video{}
	for i := range related {
		videos, err := model.GetVideos(related[i])
		if err != nil {
			return c.JSON(500, err)
		}
		rv = append(rv, videos...)
	}

	res := Response{
		Here:    hv,
		Related: rv,
	}

	return c.JSON(200, res)
}

func getFacility(longitude string, latitude string) (string, []string, error) {

	related := []string{"disney", "usj"}
	return "colony", related, nil
}
