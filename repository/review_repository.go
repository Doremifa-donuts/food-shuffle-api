package repository

import (
	"fmt"
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

// レビューを作成する
func CreateReview(db *gorm.DB, review *model.Review) error {
	return db.Create(review).Error
}

// レビューUUIDからレビューを複数取得する
func ListReviewsByReviewUuids(db *gorm.DB, reviewUuids []string) ([]model.Review, error) {
	var reviews []model.Review
	err := db.Where("review_uuid IN (?)", reviewUuids).Find(&reviews).Error
	return reviews, err
}

// レビューのステータスを更新する
func UpdateReviewStatus(db *gorm.DB, reviewFlag model.UserReviewFlag) (bool, error) {
	result := db.Model(&model.UserReviewFlag{}).Where("user_uuid = ? AND review_uuid = ?", reviewFlag.UserUuid, reviewFlag.ReviewUuid).Update("review_status", reviewFlag.ReviewStatus)
	// 更新されたレコードが1ならtrueを返却する
	fmt.Println(result.RowsAffected)
	return result.RowsAffected == 1, result.Error
}

// ユーザーUUIDをレビューUUIDの組み合わせが存在するかを確認する
func ExistReviewByUserUuidAndReviewUuid(db *gorm.DB, userUuid string, reviewUuid string) error {
	return db.Where("user_uuid = ? and review_uuid = ?", userUuid, reviewUuid).First(&model.Review{}).Error
}

func GetReviewDetail(db *gorm.DB, RestaurantUuid string, userUuid string) ([]model.Review, error) {
	var reviewDetails []model.Review
	err := db.Where("restaurant_uuid = ? and user_uuid = ?", RestaurantUuid, userUuid).Find(&reviewDetails).Error
	return reviewDetails, err
}