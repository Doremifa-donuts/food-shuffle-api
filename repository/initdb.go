package repository

import (
	"fmt"
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// DBを初期化する
func InitDB() error {
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
		return err
	}

	// GORMのロガーを設定（全てのSQLクエリを出力）
	newLogger := logger.New(
		log.New(gormLogFile, "\n", log.LstdFlags), // Writerをマルチライターで設定
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)

	// DBに接続する
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil { // DBに接続できなかった場合
		logging.LogError("Error connecting to database", err)
		return err
	} else { // DBに接続できた場合
		fmt.Println("Connected to database")
	}

	// DBのマイグレーションを実行する
	result, err := model.MigrateDB(db)
	if err != nil {
		logging.LogError("Error migrating database", err)
		return err
	} else {
		// マイグレーションが実行された場合、サンプルデータを挿入する
		if result {
			fmt.Println("Database migrated")
			// サンプルデータを挿入する
			model.InsertSampleData(db)
		} else {
			fmt.Println("Database already migrated")
		}
	}

	return nil
}

// DBを取得する
func GetDB() *gorm.DB {
	return db
}

// トランザクションを実行する関数
func Transaction(fn func(tx *gorm.DB) error) error {
	return db.Transaction(fn)
}
