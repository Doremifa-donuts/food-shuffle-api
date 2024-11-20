package repository

import (
	"gorm.io/gorm"

	"food-shuffle-api/model"
)

// ユーザーを作成する
func CreateUser(db *gorm.DB, user model.User) error {
	return db.Create(&user).Error
}

// UserUUIDが一致するユーザーのJtiTokenを更新する
func UpdateJtiTokenByUserUuid(db *gorm.DB, userUuid string, jtiToken string) error {
	return db.Model(&model.User{}).Where("user_uuid = ?", userUuid).Update("jti_token", jtiToken).Error
}

// UserUUIDとJtiTokenの組み合わせが一致するか確認
func ExistsUserByUserUuidAndJtiToken(db *gorm.DB, userUuid string, jtiToken string) error {
	var user model.User
	// UserUUIDとJtiTokenの組み合わせが一致するか確認
	err := db.Where("user_uuid = ? AND jti_token = ?", userUuid, jtiToken).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// メールアドレスが一致するユーザーを取得する
func GetUserByMailAddress(db *gorm.DB, mailAddress string) (model.User, error) {
	var user model.User
	// メールアドレスが一致するユーザーを取得する
	err := db.Where("mail_address = ?", mailAddress).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// ユーザーUUIDとUserTypeの組み合わせが一致するか確認
func ExistsUserByUserUuidAndUserType(db *gorm.DB, userUuid string, userType model.UserType) error {
	var user model.User
	err := db.Where("user_uuid = ? AND user_type = ?", userUuid, userType).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
