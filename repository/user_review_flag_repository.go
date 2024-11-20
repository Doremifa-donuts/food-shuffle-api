package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

// ステータスを限定したレビューのUUIDを取得する
func ListReviewUuidsByUserUuidAndReviewStatus(tx *gorm.DB, reviewFlag model.UserReviewFlag) ([]string, error) {
	var reviewUuids []string
	err := tx.Model(&model.UserReviewFlag{}).Select("review_uuid").Where("user_uuid = ? AND review_status = ?", reviewFlag.UserUuid, reviewFlag.ReviewStatus).Find(&reviewUuids).Error
	return reviewUuids, err
}
