package response

import "time"

// レビュー一覧のレスポンス
type GetReviewsByUser struct {
	RestaurantUuid string    // レストランUUID
	RestaurantName string    // レストラン名
	Comment        string    // レビューの内容
	PostedAt       time.Time // レビューの投稿日時
	Images         []string  // レビューの画像
	Icon           string    // ユーザーのアイコン
}
