package service

import (
	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/prefix"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewService struct{}

// すれ違いで受け取ったレビューを取得する
func (s *ReviewService) GetReceivedReviewsByUser(uuid string) ([]dto.GetReviewsByUser, error) {
	// ステータスを Unclassified に限定した構造体に変換する
	reviewFlag := model.UserReviewFlag{
		UserUuid:     uuid,
		ReviewStatus: model.Unclassified,
	}
	// 受け取ったレビューを取得する
	return s.getReviewsByUser(reviewFlag)
}

// アーカイブに登録したレビューを取得する
func (s *ReviewService) GetArchivedReviewsByUser(uuid string) ([]dto.GetReviewsByUser, error) {
	// ステータスを Interested に限定した構造体に変換する
	reviewFlag := model.UserReviewFlag{
		UserUuid:     uuid,
		ReviewStatus: model.Interested,
	}
	// アーカイブ状態のレビューを取得する
	return s.getReviewsByUser(reviewFlag)
}

// ユーザーがいいねしたレビューを取得する
func (s *ReviewService) GetLikedReviewsByUser(uuid string) ([]dto.GetReviewsByUser, error) {
	// ステータスをLikedに限定した構造体に変換する
	reviewFlag := model.UserReviewFlag{
		UserUuid:     uuid,
		ReviewStatus: model.Iiked,
	}
	// いいねしたレビューを格納するテーブルに対して、レビューを取得する
	return s.getReviewsByUser(reviewFlag)
}

// ユーザーのレビューを取得する　コールバックでレビューの種類を分ける
func (s *ReviewService) getReviewsByUser(reviewFlag model.UserReviewFlag) (res []dto.GetReviewsByUser, err error) {
	// トランザクションを開始する
	err = repository.Transaction(func(tx *gorm.DB) error {
		// ユーザーが取得するレビューのUUID
		var reviewUuids []string

		// ユーザーのレビューを取得する
		reviewUuids, err := repository.ListReviewUuidsByUserUuidAndReviewStatus(tx, reviewFlag)
		if err != nil {
			logging.LogError("failed to get reviews", err)
			return err
		}

		// レビューの構造体
		var reviews []model.Review

		// レビューの内容を取得する
		reviews, err = repository.ListReviewsByReviewUuids(tx, reviewUuids)
		if err != nil {
			logging.LogError("failed to get reviews", err)
			return err
		}

		// それぞれに不足している項目を取得する
		for _, review := range reviews {
			// レストラン名を取得する
			restaurantName, err := repository.GetRestaurantNameByRestaurantUuid(tx, review.RestaurantUuid)
			if err != nil {
				logging.LogError("failed to get restaurant name", err)
				return err
			}

			// Imagesにプレフィックスを追加する
			for i, image := range review.Images {
				review.Images[i] = prefix.ImagePrefixReview + image
			}

			// ユーザーのアイコンを取得する
			icon, err := repository.GetIconByUserUuid(tx, review.UserUuid)
			if err != nil {
				logging.LogError("failed to get icon", err)
				return err
			}

			// アイコンにプレフィックスを追加する
			icon = prefix.ImagePrefixUserIcon + icon

			// レビューをレスポンスに追加する
			res = append(res, dto.GetReviewsByUser{
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
	return
}

// ユーザーがレビューを投稿する
func (s *ReviewService) PostReview(bReview model.Review) (res dto.PostReview, err error) {
	// トランザクションを開始する
	err = repository.Transaction(func(tx *gorm.DB) error {
		// レビューUUIDを生成する
		reviewUuid, err := uuid.NewV7()
		if err != nil {
			logging.LogError("failed to generate review uuid", err)
			return err
		}

		// レビューUUIDを格納する
		bReview.ReviewUuid = reviewUuid.String()
		res.ReviewUuid = reviewUuid.String()

		// ユーザーのレビューをDBに保存する
		err = repository.CreateReview(tx, &bReview)
		if err != nil {
			logging.LogError("failed to create review", err)
			return err
		}

		// トランザクションを終了する
		return nil
	})

	// レスポンスを返却する
	return
}
