package model

import (
	"gorm.io/gorm"
)

// Article table
type Article struct {
	gorm.Model
	Title       string `json:"title" gorm:"unique"`
	Content     string `json:"content"`
	Category    uint   `json:"category"`
	Label       string `json:"label"`
	AuthorId    uint   `json:"author_id"`
	AuthorName  string `json:"author_name"`
	Description string `json:"description"`
}

// ArticlePostDTO Post
type ArticlePostDTO struct {
	Title       string          `json:"title"`
	Content     string          `json:"content"`
	Label       string          `json:"label"`
	AuthorID    uint            `json:"author_id"`
	Description string          `json:"description"`
	Category    CategoryPostDTO `json:"category"`
}

// ArticlePutDTO Get
type ArticleGetDTO struct {
	*ArticleBrief
	Content    string `json:"content"`
	AuthorName string `json:"author_name"`
}

// ArticleViewDTO Get
type ArticleBrief struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Label       string `json:"label"`
	AuthorID    uint   `json:"author_id"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
