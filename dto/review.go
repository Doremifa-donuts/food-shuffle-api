package dto

import "time"

// レビュー一覧のレスポンス
type GetReviewsByUser struct {
	RestaurantUuid string    // レストランUUID
	RestaurantName string    // レストラン名
	ReviewUuid     string    // レビューUUID
	Comment        string    // レビューの内容
	PostedAt       time.Time // レビューの投稿日時
	Images         []string  // レビューの画像
	Icon           string    // ユーザーのアイコン
}

// レビュー投稿のレスポンス
type PostReview struct {
	ReviewUuid string // レビューUUID
}

// 共有するレビュー設定のレスポンス
type ShareSettingReview struct {
	FirstShareReviewUuid  string
	SecondShareReviewUuid string
	ThirdShareReviewUuid  string
}

type ReviewDetail struct {
	ReviewUuid              string
	UserUuid                string
	RestaurantUuid          string
	Images                  []string
	CreatedAt               time.Time
	Comment                 string
}