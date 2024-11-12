package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

// 一般ユーザーの追加項目を登録する
func CreateGeneralUser(db *gorm.DB, generalUser model.GeneralUser) error {
	return db.Create(&generalUser).Error
}

// 一般ユーザーを取得する
func GetGeneralUserByUserUuid(db *gorm.DB, userUuid string) (model.GeneralUser, error) {
	var generalUser model.GeneralUser
	err := db.Where("user_uuid = ?", userUuid).First(&generalUser).Error
	if err != nil {
		return generalUser, err
	}
	return generalUser, nil
}
