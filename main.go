package main

import (
	"fmt"
	logging "food-shuffle-api/log"
	"food-shuffle-api/repository"
	"food-shuffle-api/server"
	"food-shuffle-api/utility/auth"
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

	// 認証関連のモデルを初期化する
	err = auth.InitAuth()
	if err != nil {
		fmt.Println("Error initializing authentication", err)
	}

	// ginを初期化する
	router, err := server.InitGin()
	if err != nil {
		fmt.Println("Error initializing gin", err)
	} else {
		// ポートを環境変数から取得する
		goPort := os.Getenv("GO_PORT")

		// サーバーを起動する
		router.Run(":" + goPort)
	}
}
