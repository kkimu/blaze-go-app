package model

import "time"

type Video struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
	Facility  string    `json:"facility"`
	CreatedAt time.Time `json:"created_at"`
}

func InsertVideo(video *Video) error {
	if err := db.Create(&video).Error; err != nil {
		return err
	}
	return nil
}

func UpdateVideo(video *Video) error {
	return db.Save(&video).Error
}
