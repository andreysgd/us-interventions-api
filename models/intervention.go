package models

import "gorm.io/gorm"

type Intervention struct {
	gorm.Model
	Country     string `json:"country"`
	Year        string `json:"year"`
	Description string `json:"description"`
	Motivation  string `json:"motivation"`
}
