package middleware

import (
	logging "food-shuffle-api/log"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/utility/custom_error"

	"net/http"

	"food-shuffle-api/utility/enbox"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 認証が終了するまで待機する
		// ヘッダーからトークンの文字列を取得する
		tokenString := c.GetHeader("Authorization")

		// トークンが設定されていなければエラーを返す
		if tokenString == "" {
			// エラーログを書き込む
			logging.LogError("Authorization header not found", custom_error.NewError(custom_error.ResourceNotFoundError))

			// エラーレスポンスを返す
			enbox.ResponseJson(c, http.StatusUnauthorized, nil)

			// 処理を終了する
			c.Abort()
			return
		}

		// トークンを検証する
		uuid, err := auth.ValidateToken(tokenString)
		if err != nil {
			// エラーログを書き込む
			logging.LogError("Error validating token:", err)

			// エラーレスポンスを返す
			enbox.ResponseJson(c, http.StatusUnauthorized, nil)
			// 処理を終了する
			c.Abort()
			return
		}

		// 認証に成功したUUIDをコンテキストに格納する
		c.Set("uuid", uuid)

		// 次の処理へ
		c.Next()
	}
}
