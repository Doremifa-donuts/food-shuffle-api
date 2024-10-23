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
		// ユーザー側のエンドポイントはusersグループに所属する
		users := v1.Group("/users") // v1/users
		{
			users.POST("/login", handler.LoginHandler)	// v1/users/login

			// ログイン後のエンドポイントは全てauthグループに所属する
			auth := users.Group("/auth", middleware.Auth()) // v1/users/auth
			{
				auth.POST("/")
			}
		}

		// レストラン側のエンドポイントはrestosグループに所属する
		restos := v1.Group("/restos") // v1/restos
		{
			// ログイン後のエンドポイントは全てauthグループに所属する
			auth := restos.Group("/auth", middleware.Auth()) // v1/restos/auth
			{
				auth.POST("/")
			}
		}

		// テスト用のエンドポイント
		v1.POST("/test", middleware.Auth(), func(ctx *gin.Context) {
			fmt.Println("test")
			fmt.Println(ctx.Get("uuid"))
			ctx.JSON(http.StatusOK, gin.H{"message": "test"})
		})
	}
	return router
}

// 接続確認用の静的サイトを表示する
func checkConnectionRoute(router *gin.Engine) {
	router.LoadHTMLGlob("view/*")

	router.GET("/", func(c *gin.Context) {
		httpStatus := http.StatusOK
		c.HTML(httpStatus, "index.html", gin.H{
			"status":  http.StatusText(httpStatus),
			"message": "Service is up and running!",
		})
	})

}
