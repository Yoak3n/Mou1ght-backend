package dto

import (
	"Mou1ght-Server/internal/model"
)

type ArticleDTO struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	Label       string `json:"label"`
	Author      uint   `json:"author"`
	AuthorName  string `json:"author_name"`
	Description string `json:"description"`
}

type ArticlePostDTO struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Label       string `json:"label"`
	AuthorID    uint   `json:"author_id"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type ArticleView struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Label       string `json:"label"`
	AuthorID    uint   `json:"author_id"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

func ToArticleList(as []*model.Article) []ArticleView {
	articleList := make([]ArticleView, 0)
	for _, article := range as {
		articleList = append(articleList, ArticleView{
			ID:          article.ID,
			Label:       article.Label,
			Title:       article.Title,
			AuthorID:    article.Author,
			Category:    article.Category,
			Description: article.Description,
		})
	}
	return articleList
}

func ToArticleDTO(a *model.Article) ArticleDTO {
	return ArticleDTO{
		Title:       a.Title,
		Content:     a.Content,
		Category:    a.Category,
		Label:       a.Label,
		Author:      a.Author,
		AuthorName:  a.AuthorName,
		Description: a.Description,
	}
}
