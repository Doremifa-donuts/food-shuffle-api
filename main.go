package main

import (
	"fmt"
	logging "food-shuffle-api/log"
	"food-shuffle-api/redisconn"
	"food-shuffle-api/repository"
	"food-shuffle-api/server"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/ws"
	"os"
)

func main() {
	// ログを初期化する
	err := logging.InitLogging()
	if err != nil {
		fmt.Println("Error initializing logging", err)
	}

	// DBを初期化する
	err = repository.InitDB()
	if err != nil {
		fmt.Println("Error initializing database", err)
	}

	// redis を初期化する
	redisconn.InitRedis()

	// 認証関連のモデルを初期化する
	err = auth.InitAuth()
	if err != nil {
		fmt.Println("Error initializing authentication", err)
	}

	// ウェブソケットを初期化
	ws.InitWebsocket()

	// ginを初期化する
	engine, err := server.InitGin()
	if err != nil {
		fmt.Println("Error initializing gin", err)
	} else {
		// ポートを環境変数から取得する
		goPort := os.Getenv("GO_PORT")

		// サーバーを起動する
		engine.Run(":" + goPort)
	}
}
