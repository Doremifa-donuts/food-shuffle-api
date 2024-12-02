package repository

import (
	"fmt"
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

// ユーザーのUUIDからアイコンを取得する
func GetIconByUserUuid(db *gorm.DB, userUuid string) (string, error) {
	var generalUser model.GeneralUser
	err := db.Where("user_uuid = ?", userUuid).First(&generalUser).Error
	if err != nil {
		return "", err
	}
	return generalUser.Icon, nil
}

// ユーザーUUIDのリストからステータスが通知受け取りになっている人のみに絞り込む
func ListFilterActiveStatusByUserUuids(db *gorm.DB, userUuids []string) ([]string, error) {
	var filteredUuids []string
	err := db.Model(model.GeneralUser{}).Where("user_uuid in (?) and share_status = ?", userUuids, model.Active).Pluck("user_uuid", &filteredUuids).Error
	return filteredUuids, err
}

func PutShareStatus(db *gorm.DB, generalUser model.GeneralUser) (bool, error) {
	result := db.Model(model.GeneralUser{}).Where("user_uuid = ?", generalUser.UserUuid).Update("share_status", generalUser.ShareStatus)
	//更新されたレコードが一つならtrueを返す
	fmt.Println(result.RowsAffected)
	return result.RowsAffected == 1, result.Error
}