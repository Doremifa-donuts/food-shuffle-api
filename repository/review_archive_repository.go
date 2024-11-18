package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

// アーカイブに登録したレビューを取得する
func ListArchivedReviewUuidsByUserUuid(tx *gorm.DB, uuid string) ([]string, error) {
	var reviewUuids []string
	err := tx.Model(&model.ReviewArchive{}).Where("user_uuid = ?", uuid).Select("review_uuid").Find(&reviewUuids).Error
	return reviewUuids, err
}
