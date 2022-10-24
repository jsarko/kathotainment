package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Name     string `json:"name"`
	Location string `json:"location"`
	Url      string `json:"url"`
	Img_path string `json:"image_path"`
	NoteID   int    `json:"note_id"`
}

type ActivityNote struct {
	gorm.Model
	Note       string `json:"note"`
	ActivityID int
}
