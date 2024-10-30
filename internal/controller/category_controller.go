package controller

import "Mou1ght-Server/internal/model"

func GetCategoryById(id uint) (category model.Category) {
	db.First(&category, id)
	return
}
