package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string  `json:"name"`
	Author        string  `json:"author"`
	Publisher     string  `json:"publisher"`
	PublishedDate string  `json:"published_date"`
	Rating        float64 `json:"rating"`
}
