package server

import (
	"fmt"
	"food-shuffle-api/handler"
	"food-shuffle-api/middleware"
	"food-shuffle-api/ws"
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
			// 保存した画像へのアクセスを許可　//HACK: 権限をチェックする必要がある
			auth.Static("/images", "public/images") // v1/auth/images

			// テスト用のエンドポイント
			auth.GET("/test", func(ctx *gin.Context) { // v1/auth/test
				fmt.Println("test")
				fmt.Println(ctx.Get("uuid"))
				ctx.JSON(http.StatusOK, gin.H{"message": "test"})
			})

			// お助けブースト取得
			auth.GET("/urgentCampaign/:campaignUuid", handler.GetUrgentCampaignHandler) // v1/auth/urgentcampaign

			// 店舗ごとのコース一覧取得
			auth.GET("/courses/:restaurantUuid", handler.GetCoursesHandler) // v1/auth/courses/:restaurantUuid

			// 一般ユーザー用のエンドポイント
			generals := auth.Group("/users", middleware.AllowGeneralUsers()) // v1/auth/users
			{
				// 一般ユーザーの認証をテストするエンドポイント
				generals.GET("/test", func(ctx *gin.Context) { // v1/auth/users/test
					fmt.Println("test")
					fmt.Println(ctx.Get("uuid"))
					ctx.JSON(http.StatusOK, gin.H{"message": "test"})
				})

				// WSで位置情報を送信するエンドポイント
				generals.GET("/locations", ws.LocationShareHandler) // v1/auth/users/locations

				// レビュー関連
				reviews := generals.Group("/reviews") // v1/auth/users/reviews
				{
					// すれ違いで受け取ったレビューの一覧を取得する
					reviews.GET("/recieves", handler.GetReceivedReviewsByUserHandler) // v1/auth/users/reviews/recieves

					// 興味ありに保存されたレビューの一覧を取得する
					reviews.GET("/interests", handler.GetInterestedReviewsByUserHandler) // v1/auth/users/reviews/interests

					// いいねをしたレビューの一覧を取得する
					reviews.GET("/likes", handler.GetLikedReviewsByUserHandler) // v1/auth/users/reviews/likes

					// グループにシェアしたレビューの一覧を取得する	//TODO:
					reviews.GET("/shares") // v1/auth/users/reviews/shares

					// 自分が投稿したレビューの一覧を取得する		//TODO:
					reviews.GET("/posts") // v1/auth/users/reviews/posts

					// レビューを投稿する
					reviews.POST("/post", handler.PostReviewByUserHandler) // v1/auth/users/reviews/post

					// シェアするレビューを設定する
					reviews.PUT("/set", handler.PutReviewShareSettingHandler) // v1/auth/users/reviews/set

					// レビューステータスを更新する
					reviews.PUT("/:review_uuid/status/:review_status", handler.PutReviewStatusByUserHandler) // v1/auth/users/reviews/:review_uuid/:review_status
				}

				// 一般ユーザー用のエンドポイントはこの中に追加していく

				//予約登録
				generals.POST("/reservation", handler.ReservationRegistorHandler) // v1/auth/users/reservation

				// 店舗詳細取得
				generals.GET("/restaurantDetail/:restaurantUuid", handler.GetRestaurantDetailHandler) // v1/auth/users/restaurantDetail

				// ユーザーの通知モードの変更
				generals.PUT("/putShareStatus/:status", handler.PutShareStatusHandler)	// v1/auth/users/putShareStatus
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

				restaurants.POST("/urgentCampaign", handler.CreateUrgentCampaignHandler) // v1/auth/restaurants/urgentcampaign
			}
		}

	}
	return router
}

// 接続確認用の静的サイトを表示する
func checkConnectionRoute(router *gin.Engine) {
	router.LoadHTMLGlob("public/view/*")

	router.GET("/", func(ctx *gin.Context) {
		httpStatus := http.StatusOK
		ctx.HTML(httpStatus, "index.html", gin.H{
			"status":  http.StatusText(httpStatus),
			"message": "Service is up and running!",
		})
	})

}
