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
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// 環境変数を使ってDSNを構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)

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
