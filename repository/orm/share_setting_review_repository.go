package orm

import (
	"food-shuffle-api/repository/model"

	"gorm.io/gorm"
)

// ユーザーUUIDから1件取得
func GetShareSettingReviewByUserUuid(tx *gorm.DB, userUuid string) (model.ShareSettingReview, error) {
	shareSettingReview := model.ShareSettingReview{
		UserUuid: userUuid,
	}
	err := tx.Preload("FirstReview").Where("user_uuid = ?", userUuid).First(&shareSettingReview).Error

	return shareSettingReview, err
}

// 新規作成する
func CreateShareSettingReview(db *gorm.DB, shareSettingReview model.ShareSettingReview) error {
	return db.Create(shareSettingReview).Error
}

// 更新する
func UpdateShareSettingReview(db *gorm.DB, shareSettingReview model.ShareSettingReview) error {
	return db.UpdateColumns(shareSettingReview).Error
}
