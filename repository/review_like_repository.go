package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

// ユーザーがいいねしたレビューを取得する
func ListLikedReviewUuidsByUserUuid(tx *gorm.DB, uuid string) ([]string, error) {
	var reviewUuids []string
	err := tx.Model(&model.ReviewLike{}).Where("user_uuid = ?", uuid).Select("review_uuid").Find(&reviewUuids).Error
	return reviewUuids, err
}
