package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

// RestoUUIDが一致するユーザーのJtiTokenを更新する
func SaveJtiByRestoUuid(db *gorm.DB, restoUuid string, jtiToken string) error {
	return db.Model(&model.RestoUser{}).Where("resto_uuid = ?", restoUuid).Update("jti_token", jtiToken).Error
}

// RestoUUIDとJtiTokenの組み合わせが一致するか確認
func CheckJtiResto(db *gorm.DB, restoUuid string, jtiToken string) (bool, error) {
	var restoUser model.RestoUser
	// RestoUUIDとJtiTokenの組み合わせが一致するか確認
	err := db.Where("resto_uuid = ? AND jti_token = ?", restoUuid, jtiToken).First(&restoUser).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
