package controller

import (
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
	logger.Info.Println(a.AuthorId)
	return true, result
}

// AddArticle 创建文章
func AddArticle(a *model.ArticlePostDTO) error {

	authors := make([]model.User, 0)
	user := GetUserByID(a.AuthorID)
	authors = append(authors, *user)
	updatedCategoryRecord := CreateCategory(a.Category)

	if len(authors) != 0 {
		article := &model.Article{
			Title:       a.Title,
			Description: a.Description,
			Label:       a.Label,
			Content:     a.Content,
			AuthorId:    a.AuthorID,
			AuthorName:  user.NickName,
			Category:    updatedCategoryRecord.ID,
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
func UpdateArticle(a *model.ArticlePostDTO, id uint) error {
	article := &model.Article{}
	updatedCategoryRecord := CreateCategory(a.Category)
	result := db.Model(article).Where("id = ?", id).Updates(model.Article{
		Title:       a.Title,
		Description: a.Description,
		Label:       a.Label,
		Content:     a.Content,
		AuthorId:    a.AuthorID,
		Category:    updatedCategoryRecord.ID,
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

// GetArticlesByAuthorId 通过作者ID获取文章
func GetArticlesByAuthorId(id uint) ([]*model.Article, error) {
	articles := make([]*model.Article, 0)
	result := db.Where("author_id = ?", id).Find(articles)
	if result.RowsAffected == 0 {
		return nil, errors.New("article not found")
	}
	return articles, nil
}
