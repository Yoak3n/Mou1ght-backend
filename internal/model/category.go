package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
	ParentID    uint
}
type CategoryPostDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentID    uint   `json:"parent_id"`
}
