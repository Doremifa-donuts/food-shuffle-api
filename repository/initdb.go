package repository

import (
	"fmt"
	logging "food-shuffle-api/log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	// コンテナに設定されている環境変数を読み込む
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")

	// 環境変数を使ってDSNを構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE)

	// DBに接続する
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil { // DBに接続できなかった場合
		logging.LogError("Error connecting to database", err)
	} else { // DBに接続できた場合
		fmt.Println("Connected to database")
	}
	return db
}

func GetDB() *gorm.DB {
	return db
}
