package dto

// レストランの詳細情報を取得したときのレスポンス
type RestaurantDetail struct {
	RestaurantUuid string
	RestaurantName string
	Address        string
	Images         []string
	Url            string
	Summary        string
	BusinessHours  string
}

type VisitedRestaurants struct {
	RestaurantUuid string
	RestaurantName string
	Address        string
	Images         []string
	Url            string
	Summary        string
	BusinessHours  string
}


// 位置情報のリクエスト
type CheckInLocation struct {
	Location struct {
		Latitude  float64 `binding:"required"`
		Longitude float64 `binding:"required"`
	}
}
