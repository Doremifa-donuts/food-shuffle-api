package service

import (
	"errors"
	"food-shuffle-api/dto/response"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/custom_error"
	"food-shuffle-api/utility/prefix"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewService struct{}

// すれ違いで受け取ったレビューを取得する
func (s *ReviewService) GetReceivedReviewsByUser(uuid string) ([]response.GetReviewsByUser, error) {
	// 受け取ったレビューを格納するテーブルに対して、レビューを取得する
	return s.getReviewsByUser(uuid, repository.ListReceivedReviewUuidsByUserUuid)
}

// アーカイブに登録したレビューを取得する
func (s *ReviewService) GetArchivedReviewsByUser(uuid string) ([]response.GetReviewsByUser, error) {
	// アーカイブに登録したレビューを格納するテーブルに対して、レビューを取得する
	return s.getReviewsByUser(uuid, repository.ListArchivedReviewUuidsByUserUuid)
}

// ユーザーがいいねしたレビューを取得する
func (s *ReviewService) GetLikedReviewsByUser(uuid string) ([]response.GetReviewsByUser, error) {
	// いいねしたレビューを格納するテーブルに対して、レビューを取得する
	return s.getReviewsByUser(uuid, repository.ListLikedReviewUuidsByUserUuid)
}

// ユーザーのレビューを取得する　コールバックでレビューの種類を分ける
func (s *ReviewService) getReviewsByUser(uuid string, callback func(tx *gorm.DB, uuid string) ([]string, error)) ([]response.GetReviewsByUser, error) {
	// レスポンスの型を定義する
	var res []response.GetReviewsByUser
	// ユーザーのレビューを取得する

	err := repository.Transaction(func(tx *gorm.DB) error {

		// ユーザーが取得するレビューのUUID
		var reviewUuids []string
		// ユーザーのレビューを取得する
		reviewUuids, err := callback(tx, uuid)
		if err != nil {
			return err
		}

		// レビューの構造体
		var reviews []model.Review

		// レビューの内容を取得する
		reviews, err = repository.ListReviewsByReviewUuids(tx, reviewUuids)
		if err != nil {
			return err
		}

		// それぞれに不足している項目を取得する
		for _, review := range reviews {
			// レストラン名を取得する
			restaurantName, err := repository.GetRestaurantNameByRestaurantUuid(tx, review.RestaurantUuid)
			if err != nil {
				return err
			}

			// Imagesにプレフィックスを追加する
			for i, image := range review.Images {
				review.Images[i] = prefix.ImagePrefixReview + image
			}

			// ユーザーのアイコンを取得する
			icon, err := repository.GetIconByUserUuid(tx, review.UserUuid)
			if err != nil {
				return err
			}

			// アイコンにプレフィックスを追加する
			icon = prefix.ImagePrefixUserIcon + icon

			// レビューをレスポンスに追加する
			res = append(res, response.GetReviewsByUser{
				RestaurantUuid: review.RestaurantUuid,
				RestaurantName: restaurantName,
				Comment:        review.Comment,
				PostedAt:       review.CreatedAt,
				Images:         review.Images,
				Icon:           icon,
			})
		}

		// エラーが出なかった場合はnilでトランザクションを終了する
		return nil
	})

	// レスポンスを返却する
	return res, err
}

// ユーザーがレビューを投稿する
func (s *ReviewService) PostReview(review model.Review) (string, error) {

	// トランザクションを開始する
	err := repository.Transaction(func(tx *gorm.DB) error {

		// レストランに訪れたことがあるかを確認する
		err := repository.ExistsUserVisitedRestaurantByUserUuid(tx, review.UserUuid)
		if err != nil {
			// レストランに訪れたことがない場合
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return custom_error.NewError(http.StatusForbidden, "You have not visited the restaurant")
			}
			return err
		}
		// レビューUUIDを生成する
		reviewUuid, err := uuid.NewV7()
		if err != nil {
			return err
		}
		// レビューUUIDを追加する
		review.ReviewUuid = reviewUuid.String()

		// ユーザーのレビューをDBに保存する
		err = repository.CreateReview(tx, &review)
		if err != nil {
			return err
		}
		return nil
	})

	return review.ReviewUuid, err

}
