package dto

import "time"

// レビュー一覧のレスポンス
type GetReviewsByUser struct {
	RestaurantUuid string      // レストランUUID
	RestaurantName string      // レストラン名
	ReviewUuid     string      // レビューUUID
	Comment        string      // レビューの内容
	CreatedAt      time.Time   // レビューの投稿日時
	Images         []string    // レビューの画像
	Icon           string      // ユーザーのアイコン
	Good           int         // いいねの数
	Address        string      // 店舗の住所
	Geolocation    Geolocation //店舗の位置情報
}
type Geolocation struct {
	Latitude  float32
	Longitude float32
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

// レビューの詳細
type ReviewDetail struct {
	ReviewUuid     string
	UserUuid       string
	RestaurantUuid string
	Images         []string
	CreatedAt      time.Time
	Comment        string
}

// 特定の店舗に対するレビューの一覧
type SpecificReviews struct {
	ReviewUuid string
	Comment    string    // レビューの内容
	CreatedAt  time.Time // レビューの投稿日時
	Images     []string  // レビューの画像
	Icon       string    // ユーザーのアイコン
	Good       int       // いいねの数
	GoodFlag   bool
}
