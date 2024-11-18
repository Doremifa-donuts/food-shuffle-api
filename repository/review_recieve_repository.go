package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

// 自身が受け取ったレビューを取得する
func ListReceivedReviewUuidsByUserUuid(db *gorm.DB, userUuid string) ([]string, error) {
	var reviewUuids []string
	err := db.Model(&model.ReviewReceive{}).Where("user_uuid = ?", userUuid).Select("review_uuid").Find(&reviewUuids).Error
	return reviewUuids, err
}
