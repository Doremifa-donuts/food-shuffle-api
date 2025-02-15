package dto

type BusyStatus string
// レストランの詳細情報を取得したときのレスポンス
const (
	Free BusyStatus = "Free"
	Spare BusyStatus = "Spare"
	Packed BusyStatus = "Packed"
)
type RestaurantDetail struct {
	RestaurantUuid string
	RestaurantName string
	Address        string
	Tell           string
	Images         []string
	Url            string
	Summary        string
	BusinessHours  string
	BusyStatus     BusyStatus
}

// 位置情報のリクエスト
type CheckInLocation struct {
	Location struct {
		Latitude  float64 `binding:"required"`
		Longitude float64 `binding:"required"`
	}
}

type WentPlaces struct {
	RestaurantUuid string  `json:"restaurant_uuid" gorm:"column:restaurant_uuid"`
	RestaurantName string  `json:"restaurant_name" gorm:"column:restaurant_name"`
	Latitude       float64 `json:"latitude" gorm:"column:latitude"`
	Longitude      float64 `json:"longitude" gorm:"column:longitude"`
}

// 一般ユーザーの仮登録に使う構造体
type PreRegisterRequest struct {
	MailAddress     string
	UserName        string
	Password        string
	ConfirmPassword string
	Tell            string
}

// 一般ユーザーの仮登録のレスポンス
type PreRegisterResponse struct {
	Key string
}

// 一般ユーザーの本登録に使う構造体
type RegisterRequest struct {
	PreRegisterKey string
	Token          string
}
