package model

import (
	"time"
)

// UserVisitedRestaurant ユーザーが行ったレストランを管理するテーブル
type UserVisitedRestaurant struct {
	UserUuid       string    `gorm:"type:char(36);primary_key"`
	RestaurantUuid string    `gorm:"type:char(36);primary_key"`
	CreatedAt      time.Time `gorm:"not null"`
}

// サンプルデータ
var sampleUserVisitedRestaurants = []UserVisitedRestaurant{
	{
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
	},
	{
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "a80499ae-eb6c-1305-a5cc-e1510c52744a",
	},
}
