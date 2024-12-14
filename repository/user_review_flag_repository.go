package repository

import (
	"food-shuffle-api/model"

	"food-shuffle-api/dto"

	"gorm.io/gorm"
)

// ステータスを限定したレビューのUUIDを取得する
func ListReviewUuidsByUserUuidAndReviewStatus(tx *gorm.DB, reviewFlag model.UserReviewFlag) ([]string, error) {
	var reviewUuids []string
	err := tx.Model(&model.UserReviewFlag{}).Select("review_uuid").Where("user_uuid = ? AND review_status = ?", reviewFlag.UserUuid, reviewFlag.ReviewStatus).Find(&reviewUuids).Error
	return reviewUuids, err
}

func ListReviewUuidsByUserUuidUnClassified(tx *gorm.DB, userUuid string) ([]dto.UnclassifiedReview, error) {
	var unclassifiedReviews []dto.UnclassifiedReview

	err := tx.Model(&model.Review{}).
		Select("reviews.*,restaurant_users.address, (SELECT COUNT(*) FROM user_review_flags WHERE review_uuid = reviews.review_uuid AND review_status = ?) AS goods", model.Iiked).
		Joins("LEFT JOIN user_review_flags AS urf ON reviews.review_uuid = urf.review_uuid AND urf.user_uuid = ?", userUuid).
		Joins("LEFT JOIN restaurant_users ON reviews.restaurant_uuid = restaurant_users.restaurant_uuid").
		Where("urf.review_uuid IS NULL OR urf.review_status = ?", model.Unclassified).
		Where("reviews.user_uuid != ?", userUuid).
		Find(&unclassifiedReviews).
		Error

	return unclassifiedReviews, err
}

// 特定のレビューに対するいいね数をカウントする
func CountReviewLikesByReviewUuid(tx *gorm.DB, reviewUuid string) (int64, error) {
	var count int64
	err := tx.Model(&model.UserReviewFlag{}).Where("review_status = ? and review_uuid = ?", model.Iiked, reviewUuid).Count(&count).Error
	return count, err
}

// レビューの受け取り処理を行う
func CreateUserReviewFlag(db *gorm.DB, reviewUuid string, userUuids []string) error {
	// 構造体にバインドする
	var userList []model.UserReviewFlag
	for _, userUuid := range userUuids {
		user := model.UserReviewFlag{
			ReviewUuid: reviewUuid,
			UserUuid:   userUuid,
		}
		userList = append(userList, user)
	}
	err := db.Create(userList).Error
	return err
}

// レビューを所持していない人のみに絞り込む
func ListExcludeUserUuidByReviewUuid(db *gorm.DB, reviewUuid string, userUuids []string) ([]string, error) {
	var excludedUserUuids []string

	// サブクエリを構築
	subQuery := db.Model(&model.UserReviewFlag{}).Select("user_uuid").
		Where("review_uuid = ?", reviewUuid).
		Where("user_uuid IN ?", userUuids)

	// メインクエリでNOT INを使用
	err := db.Model(&model.User{}).
		Where("user_uuid IN ?", userUuids).
		Where("user_uuid NOT IN (?)", subQuery).
		Pluck("user_uuid", &excludedUserUuids).Error

	return excludedUserUuids, err
}
