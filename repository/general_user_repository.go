package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

// 一般ユーザーの追加項目を登録する
func CreateGeneralUser(db *gorm.DB, generalUser model.GeneralUser) error {
	return db.Create(&generalUser).Error
}
