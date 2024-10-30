package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string `json:"title" gorm:"unique"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	Label       string `json:"label"`
	Author      uint   `json:"author"`
	AuthorName  string `json:"author_name"`
	Description string `json:"description"`
}
