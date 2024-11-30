package dto

//レストランの詳細情報を取得したときのレスポンス
type RestaurantDetail struct {
	RestaurantUuid string
	RestaurantName string
	Address        string
	Images         []string
	Url            string
	Summary        string
	BusinessHours  string
}