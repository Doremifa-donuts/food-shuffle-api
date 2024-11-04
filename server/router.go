package server

import (
	"fmt"
	"food-shuffle-api/handler"
	"food-shuffle-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routing(router *gin.Engine) *gin.Engine {
	// エンドポイントのURLは「/」区切りでグループに所属する
	v1 := router.Group("/v1") // http://IPADDRESS:5678/v1/
	{
		v1.POST("/login", handler.LoginHandler) // v1/login

		v1.POST("/users/register", handler.GeneralUserRegisterHandler) // v1/users/register

		// ログイン後のエンドポイントは全てauthグループに所属する
		auth := v1.Group("/auth", middleware.Auth()) // v1/auth/
		{
			// テスト用のエンドポイント
			v1.GET("/test", func(ctx *gin.Context) { // v1/auth/test
				fmt.Println("test")
				fmt.Println(ctx.Get("uuid"))
				ctx.JSON(http.StatusOK, gin.H{"message": "test"})
			})

			// 一般ユーザー用のエンドポイント
			generals := auth.Group("/users", middleware.AllowGeneralUsers()) // v1/auth/users
			{
				// 一般ユーザーの認証をテストするエンドポイント
				generals.GET("/test", func(ctx *gin.Context) { // v1/auth/users/test
					fmt.Println("test")
					fmt.Println(ctx.Get("uuid"))
					ctx.JSON(http.StatusOK, gin.H{"message": "test"})
				})

				// 一般ユーザーのアカウント作成
				// generals.POST("/register", handler.GeneralUserRegisterHandler)	// v1/auth/users/register

				// 一般ユーザー用のエンドポイントはこの中に追加していく

			}

			// レストランユーザー用のエンドポイント
			restaurants := auth.Group("/restaurants", middleware.AllowRestaurantUsers()) // v1/auth/restorants
			{
				// レストランユーザーの認証をテストするエンドポイント
				restaurants.GET("/test", func(ctx *gin.Context) {
					fmt.Println("test")
					fmt.Println(ctx.Get("uuid"))
					ctx.JSON(http.StatusOK, gin.H{"message": "test"})
				})

				// レストラン用のエンドポイントはこの中に追加していく
				reservations := restaurants.Group("/reservations")
				{
					// 予約の一覧を取得する
					reservations.GET("/", handler.GetReservationsHandler) // v1/auth/restaurants/reservations/
				}
			}
		}

	}
	return router
}

// 接続確認用の静的サイトを表示する
func checkConnectionRoute(router *gin.Engine) {
	router.LoadHTMLGlob("view/*")

	router.GET("/", func(ctx *gin.Context) {
		httpStatus := http.StatusOK
		ctx.HTML(httpStatus, "index.html", gin.H{
			"status":  http.StatusText(httpStatus),
			"message": "Service is up and running!",
		})
	})

}
