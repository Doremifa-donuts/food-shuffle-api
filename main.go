package main

import (
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/server"
	"food-shuffle-api/utility/auth"
	"os"
)

func main() {
	// ログを初期化する
	logging.InitLogging()

	// DBを初期化する
	repository.InitDB()

	// モデルを初期化する
	model.MigrateDB(repository.GetDB())

	// 認証関連のモデルを初期化する
	auth.InitAuth()

	// ginを初期化する
	router := server.InitGin()

	goPort := os.Getenv("GO_PORT")

	// サーバーを起動する
	router.Run(":" + goPort)
}
