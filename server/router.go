package server

import (
	"food-shuffle-api/handler"
	"food-shuffle-api/middleware"
	"food-shuffle-api/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routing(router *gin.Engine) *gin.Engine {
	// APIバージョン1グループ 将来の拡張性を考えて
	v1 := router.Group("/v1") // http://IPADDRESS:5678/v1/
	{

		v1.POST("/login", handler.LoginHandler) // v1/login

		v1.POST("/register", handler.GeneralUserRegisterHandler) // v1/register

		// ログイン後のエンドポイントは全てauthグループに所属する
		auth := v1.Group("/auth", middleware.Auth()) // v1/auth/
		{
			// 画像取得エンドポイント
			auth.GET("/images/:image_id", handler.GetImagesHandler)

			// お助けブースト1件取得　//FIXME: 店舗とユーザー側でエンドポイントを分割する
			auth.GET("/campaigns/:campaign_uuid", handler.GetUrgentCampaignHandler) // v1/auth/campaigns/canpaigns_uuid

			// 店舗ごとのコース一覧取得	//FIXME: 店舗とユーザー側でエンドポイントを分割する
			auth.GET("/courses/:restaurant_uuid", handler.GetCoursesHandler) // v1/auth/courses/:restaurant_uuid)

			// 一般ユーザー用のエンドポイント
			generals := auth.Group("/users", middleware.AllowGeneralUsers()) // v1/auth/users
			{
				// ユーザーの行ったところの情報を取得する
				generals.GET("/places", handler.GetWentPlacesHandler) // v1/auth/users/

				// WSで位置情報を送信するエンドポイント
				generals.GET("/locations", ws.LocationShareHandler) // v1/auth/users/locations

				// ユーザーの通知モードの変更
				generals.PUT("/mode/:status", handler.PutShareStatusHandler) // v1/auth/users/mode/:status

				// レビュー関連
				reviews := generals.Group("/reviews") // v1/auth/users/reviews
				{
					// すれ違いで受け取ったレビューの一覧を取得する	// FIXME: これきもい
					reviews.GET("/recieves", handler.GetReceivedReviewsByUserHandler) // v1/auth/users/reviews/recieves

					// 興味ありに保存されたレビューの一覧を取得する　// FIXME: これきもい
					reviews.GET("/interests", handler.GetInterestedReviewsByUserHandler) // v1/auth/users/reviews/interests

					// いいねをしたレビューの一覧を取得する　// FIXME: これきもい
					reviews.GET("/likes", handler.GetLikedReviewsByUserHandler) // v1/auth/users/reviews/likes

					// グループにシェアしたレビューの一覧を取得する	//TODO:
					reviews.GET("/shares") // v1/auth/users/reviews/shares

					// 自分が投稿したレビューの一覧を取得する		//TODO:
					reviews.GET("/posts") // v1/auth/users/reviews/posts

					// レビューを投稿する
					reviews.POST("/posts", handler.PostReviewByUserHandler) // v1/auth/users/reviews/posts

					// シェアするレビューを設定する //FIXME: レビューを複数件同時にシェアする方法を考える
					reviews.PUT("/sets", handler.PutReviewShareSettingHandler) // v1/auth/users/reviews/sets

					// レビューステータスを更新する
					reviews.PUT("/:review_uuid/status/:review_status", handler.PutReviewStatusByUserHandler) // v1/auth/users/reviews/:review_uuid/:review_status
				}

				// 予約関連
				reservations := generals.Group("/reservations")
				{
					reservations.GET("/upcomings", handler.GetUserUpcomingsReservationsHandler)
				}

				// ユーザーがレストランに付随する情報を操作するグループ
				restaurants := generals.Group("/restaurants")
				{
					// 行った店の店舗情報一覧取得
					restaurants.GET("/visited", handler.GetVisitedRestaurantsHandler) // v1/auth/users/visitedRestaurants
					restaurants.GET("/reviewed", handler.GetReviewedRestaurantsHandler)

					// 1つの店舗に対する情報を取得するグループ
					info := restaurants.Group("/:restaurant_uuid")
					{
						// 店舗詳細取得
						info.GET("/", handler.GetRestaurantDetailHandler) // v1/auth/users/restaurantDetail

						//予約登録
						info.POST("/reservations", handler.ReservationRegisterHandler) // v1/auth/users/restaurants/:restaurant_uuid/reservations

						// 自分が投稿したレビューを取得する
						info.GET("/posted", handler.GetPostedReviewHandler) // v1/auth/users/restaurants/:restaurant_uuid/posted

						// 自身が受け取った店舗に対するレビューを取得する
						info.GET("/reviews", handler.GetSpecificRestaurantReviewHandler)
						// 店舗へのチェックインを行うエンドポイント
						info.POST("/checkin", handler.PostCheckinRestaurantHandler)
					}
				}
			}

			// レストランユーザー用のエンドポイント
			restaurants := auth.Group("/restaurants", middleware.AllowRestaurantUsers()) // v1/auth/restorants
			{
				// 自身の店舗情報を取得する
				restaurants.GET("/", handler.GetOwnRestaurantDetailHandler)

				restaurants.GET("/courses", handler.GetOwnCoursesHandler)

				//レビュー情報取得
				restaurants.GET("/reviews", handler.GetOwnReviewsHandler) // v1/auth/restaurants/reviews

				// レストラン用のエンドポイントはこの中に追加していく
				reservations := restaurants.Group("/reservations")
				{
					// 予約の一覧を取得する
					reservations.GET("/", handler.GetReservationsHandler) // v1/auth/restaurants/reservations/
				}

				// お助けブーストの作成
				restaurants.POST("/campaigns", handler.CreateUrgentCampaignHandler) // v1/auth/restaurants/campaigns

				//混雑状況の切り替え
				restaurants.PUT("/mode/:status", handler.PutBusyStatusHandler) // v1/auth/restaurants/mode/:status
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
