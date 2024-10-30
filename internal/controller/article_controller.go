package controller

import (
	"Mou1ght-Server/internal/dto"
	"Mou1ght-Server/internal/model"
	"Mou1ght-Server/package/logger"
	"errors"

	"gorm.io/gorm"
)

func CheckExistArticle(a *model.Article, id uint) (bool, *gorm.DB) {
	result := db.First(a, id)
	if result.RowsAffected == 0 {
		return false, nil
	}
	logger.Info.Println(a.Author)
	return true, result
}

// AddArticle 创建文章
func AddArticle(a *dto.ArticlePostDTO) error {

	authors := make([]model.User, 0)
	user := GetUserByID(a.AuthorID)
	authors = append(authors, *user)

	if len(authors) != 0 {
		article := &model.Article{
			Title:       a.Title,
			Description: a.Description,
			Label:       a.Label,
			Content:     a.Content,
			Author:      a.AuthorID,
			AuthorName:  user.NickName,
			Category:    a.Category,
		}
		result := db.Create(article)
		if result.Error != nil {
			return result.Error
		}
		logger.Info.Println("Add article successfully")
		return nil
	} else {
		return errors.New("unauthorized user")
	}

}

// DeleteArticle 删除文章
func DeleteArticle(id uint) error {
	article := &model.Article{}
	result := db.Delete(article, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateArticle 更新文章
func UpdateArticle(a *dto.ArticlePostDTO, id uint) error {
	article := &model.Article{}
	result := db.Model(article).Where("id = ?", id).Updates(model.Article{
		Title:       a.Title,
		Description: a.Description,
		Label:       a.Label,
		Content:     a.Content,
		Author:      a.AuthorID,
		Category:    a.Category,
	})
	return result.Error
}

// GetArticleList 获取文章列表
func GetArticleList() ([]*model.Article, error) {

	articles := make([]*model.Article, 0)
	result := db.Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

// GetArticleById 通过ID获取文章
func GetArticleById(id uint) (*model.Article, error) {
	article := &model.Article{}
	result := db.First(article, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("article not found")
	}
	return article, nil
}

// GetArticlesByLabel 通过标题获取文章
func GetArticlesByTitle(label string) ([]*model.Article, error) {
	articles := make([]*model.Article, 0)
	result := db.Where("label = ?", label).Find(articles)
	if result.RowsAffected == 0 {
		return nil, errors.New("article not found")
	}
	return articles, nil
}
