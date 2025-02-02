package dto

import "time"

// 予約一覧のレスポンスの構造体
type ReservationsByRestaurant struct {
	ReservationUuid   string    // 予約UUID
	ReservationDate   time.Time // 予約日
	NumberOfPeople    int       // 予約人数
	UserName          string    // ユーザー名
	ReservationStatus bool      // 予約が通ってるかどうかのフラグ
}

// 予約登録時のレスポンスの構造体
type PostReservation struct {
	ReservationUuid string // 予約UUID
}

// ユーザーが予約状況を取得した場合のレスポンス
type UserReservation struct {
	RestaurantUuid    string
	RestaurantName    string
	ReservationDate   time.Time
	NumberOfPeople    int
	// CourseUuid        string
	// CourseName        string
	// CampaignUuid      string
	// CampaignName      string
	// Price             int
	ReservationStatus bool
}
