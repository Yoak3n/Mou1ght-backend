package dto

import (
	"Mou1ght-Server/internal/controller"
	"Mou1ght-Server/internal/model"
)

func ToArticleList(as []*model.Article) []model.ArticleBrief {
	articleList := make([]model.ArticleBrief, 0)
	for _, article := range as {
		ao := ToArticleDTO(article)
		av := *ao.ArticleBrief
		articleList = append(articleList, av)
	}
	return articleList
}

func ToArticleDTO(a *model.Article) model.ArticleGetDTO {
	categoryName := ""
	if a.Category != 0 {
		categoryName = controller.GetCategoryById(a.Category).Name
	}
	return model.ArticleGetDTO{
		ArticleBrief: &model.ArticleBrief{
			ID:          a.ID,
			Title:       a.Title,
			Label:       a.Label,
			AuthorID:    a.AuthorId,
			Description: a.Description,
			Category:    categoryName,
		},
		Content:    a.Content,
		AuthorName: a.AuthorName,
	}
}
