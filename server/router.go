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
			// 保存した画像へのアクセスを許可　//HACK: 権限をチェックする必要がある
			auth.Static("/images", "public/images") // v1/auth/images

			// テスト用のエンドポイント
			auth.GET("/test", func(ctx *gin.Context) { // v1/auth/test
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

				// レビュー関連
				reviews := generals.Group("/reviews") // v1/auth/users/reviews
				{
					// すれ違いで受け取ったレビューの一覧を取得する
					reviews.GET("/recieves", handler.GetReceivedReviewsByUserHandler) // v1/auth/users/reviews/recieves

					// アーカイブに保存されたレビューの一覧を取得する
					reviews.GET("/archives", handler.GetArchivedReviewsByUserHandler) // v1/auth/users/reviews/archives

					// いいねをしたレビューの一覧を取得する
					reviews.GET("/likes", handler.GetLikedReviewsByUserHandler) // v1/auth/users/reviews/likes

					// グループにシェアしたレビューの一覧を取得する	//TODO:
					reviews.GET("/shares") // v1/auth/users/reviews/shares

					// 自分が投稿したレビューの一覧を取得する		//TODO:
					reviews.GET("/posts") // v1/auth/users/reviews/posts

					// レビューを投稿する	//TODO: 画像の保存より前に投稿権限があるかを確認する
					reviews.POST("/post", middleware.AllowReviewPost(),handler.PostReviewByUserHandler) // v1/auth/users/reviews/post

				}

				// 一般ユーザー用のエンドポイントはこの中に追加していく

				generals.POST("/reservation", handler.ReservationRegistorHandler) // v1/auth/users/reservation
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
	router.LoadHTMLGlob("public/view/*")

	router.GET("/", func(ctx *gin.Context) {
		httpStatus := http.StatusOK
		ctx.HTML(httpStatus, "index.html", gin.H{
			"status":  http.StatusText(httpStatus),
			"message": "Service is up and running!",
		})
	})

}
