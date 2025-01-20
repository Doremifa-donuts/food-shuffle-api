package model

import (
	"time"
)

// UserVisitedRestaurant ユーザーが行ったレストランを管理するテーブル
type UserVisitedRestaurant struct {
	UserUuid       string    `gorm:"type:char(36);primary_key"`
	RestaurantUuid string    `gorm:"type:char(36);primary_key"`
	CreatedAt      time.Time `gorm:"not null"`
	UpdatedAt      time.Time `gorm:"not null"`
}

// サンプルデータ
var sampleUserVisitedRestaurants = []UserVisitedRestaurant{
	{ // 路地裏ビストロ816
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
	},
	{ // 炭焼 高田屋
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "a80499ae-eb6c-1305-a5cc-e1510c52744a",
	},
	{ // おにぎりごりちゃん 中崎町本店
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "0bf97fc8-019e-421b-85f5-84818aab19d8",
	},
	{ // Cafe de paris 大丸心斎橋店
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "d61aed9f-68b0-4efd-af77-98d7e061526d",
	},
	{ // かごのや
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "6d7c3625-a1fa-4d63-8600-39f538dcac87",
	},
	{ // ラーメン大王 西中島店
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "5923b6b8-a4d6-4419-acf1-b1410480b0b5",
	},
}
