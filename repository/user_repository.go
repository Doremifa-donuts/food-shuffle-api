package repository

import (
	"gorm.io/gorm"

	"food-shuffle-api/model"
)

// UserUUIDが一致するユーザーのJtiTokenを更新する
func SaveJtiByUserUuid(db *gorm.DB, userUuid string, jtiToken string) error {
	return db.Model(&model.User{}).Where("user_uuid = ?", userUuid).Update("jti_token", jtiToken).Error
}

// UserUUIDとJtiTokenの組み合わせが一致するか確認
func CheckJtiUser(db *gorm.DB, userUuid string, jtiToken string) (bool, error) {
	var user model.User
	// UserUUIDとJtiTokenの組み合わせが一致するか確認
	err := db.Where("user_uuid = ? AND jti_token = ?", userUuid, jtiToken).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
