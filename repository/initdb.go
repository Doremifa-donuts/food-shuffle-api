package repository

import (
	"fmt"
	logging "food-shuffle-api/log"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() {
	// コンテナに設定されている環境変数を読み込む
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// 環境変数を使ってDSNを構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)

	// GORMのログファイルを設定
	gormLogFile, err := os.OpenFile("log/gorm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logging.LogError("Error opening gorm.log", err)
	}

	// GORMのロガーを設定（全てのSQLクエリを出力）
	newLogger := logger.New(
		log.New(io.MultiWriter(gormLogFile, os.Stdout), "\r\n", log.LstdFlags), // Writerをマルチライターで設定
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)

	// DBに接続する
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{ Logger: newLogger })
	if err != nil { // DBに接続できなかった場合
		logging.LogError("Error connecting to database", err)
	} else { // DBに接続できた場合
		fmt.Println("Connected to database")
	}


}

func GetDB() *gorm.DB {

	return db
}
