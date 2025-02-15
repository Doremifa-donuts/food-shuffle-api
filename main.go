package main

import (
	"fmt"
	logging "food-shuffle-api/log"
	"food-shuffle-api/repository/fireauth"
	"food-shuffle-api/repository/orm"
	"food-shuffle-api/repository/redis"
	"food-shuffle-api/server"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/utility/cron"
	"food-shuffle-api/ws"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// ログを初期化する
	err := logging.InitLogging()
	if err != nil {
		fmt.Println("Error initializing logging", err)
	}

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// DBを初期化する
	err = orm.InitDB()
	if err != nil {
		fmt.Println("Error initializing database", err)
	}

	// redis を初期化する
	redis.InitRedis()

	// 認証関連のモデルを初期化する
	err = auth.InitAuth()
	if err != nil {
		fmt.Println("Error initializing authentication", err)
	}

	// firebase adminSDK の初期化
	err = fireauth.InitFirebase()
	if err != nil {
		fmt.Println("Error initializing authentication", err)
	}

	// ウェブソケットを初期化
	ws.InitWebsocket()

	// タスクスケジューラーを起動
	cron.Run()

	// ginを初期化する
	engine, err := server.InitGin()
	if err != nil {
		fmt.Println("Error initializing gin", err)
	} else {
		// ポートを環境変数から取得する
		goPort := os.Getenv("GO_PORT")

		// サーバーをgoroutineで実行
		engine.Run(":" + goPort)
	}
}
