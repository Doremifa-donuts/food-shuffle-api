package ws

type Status int

const (
	Review Status = iota // 0
	Boost                // 1
)

// お助けブーストの通知内容
type BoostContent struct {
	RestaurantName string
	BoostUuid      string
}

// レビュー受信の通知内容
type ReviewContent struct {
}

// 通知の対象と通知内容を持つ構造体
type NotifyProvider struct {
	UserUuids []string `json:"-"`
	Type      Status
	Message   string
	Content   interface{}
}
