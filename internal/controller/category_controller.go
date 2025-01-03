package controller

import (
	"Mou1ght-Server/internal/model"

	"gorm.io/gorm/clause"
)

func CreateCategory(category model.CategoryPostDTO) model.Category {
	categoryRecord := model.Category{
		Name:        category.Name,
		Description: category.Description,
	}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "description"}), // columns to update
	}).Create(&categoryRecord)
	// 是否会自动更新id？
	return categoryRecord
}

func GetCategoryById(id uint) (category model.Category) {
	db.First(&category, id)
	return
}
