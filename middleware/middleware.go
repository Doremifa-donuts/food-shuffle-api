package middleware

import (
	"errors"
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/utility/custom_error"

	"net/http"

	"food-shuffle-api/utility/conversion"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// ユーザーの認証を行い、対象のuuidをコンテキストに格納する
func Auth() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// 認証が終了するまで待機する
		// ヘッダーからトークンの文字列を取得する
		tokenString := ctx.GetHeader("Authorization")

		// トークンが設定されていなければエラーを返す
		if tokenString == "" {
			// エラーログを書き込む
			logging.LogError("Authorization header not found", nil)

			// エラーレスポンスを返す
			conversion.ResponseJson(ctx, http.StatusUnauthorized, nil)

			// 処理を終了する
			ctx.Abort()
			return
		}

		// トークンを検証する
		uuid, err := auth.ValidateToken(tokenString)
		if err != nil {
			// エラーログを書き込む
			logging.LogError("Error validating token:", err)

			// エラーによって異なるレスポンスを返す
			// カスタムエラーの変数を宣言
			var customErr *custom_error.CustomError

			// カスタムエラーにキャスト可能か確認する
			// キャスト可能な場合、自身で定義したビジネスロジック上のエラーなので、適切なレスポンスを返す
			if errors.As(err, &customErr) {
				// エラー文は発生した時点でログに書き込んだことが良い気がするのでレスポンス分けたいエラー以外はそこまで詳しく分類しなくていいかもしれない
				conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
				ctx.Abort()
				return
			}

			// mysqlのエラー
			// mysqlエラー型の定義
			var mySQLError *mysql.MySQLError
			if errors.As(err, &mySQLError) {
				logging.LogError("mysql error", err)
				conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
				ctx.Abort()
				return
			}

			// その他のエラー
			conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
			// 処理を終了する
			ctx.Abort()
			return
		}

		// 認証に成功したUUIDをコンテキストに格納する
		ctx.Set("uuid", uuid)

		// 次の処理へ
		ctx.Next()
	}
}

// ユーザーのアカウントが適切かを判定する
func authorizeUserType(ctx *gin.Context, userType model.UserType) {
	// idが存在しなければエラーを返す
	uuid, _ := ctx.Get("uuid")
	// 型アサーション
	uuidAdjusted := uuid.(string)

	// ユーザーのアカウントタイプをチェック
	err := repository.ExistsUserByUserUuidAndUserType(repository.GetDB(), uuidAdjusted, userType)
	if err != nil {
		// エラーを分類する
		// リソースが見つからない場合は、権限のないエンドポイントへの通信を行ったことを意味する
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// エラーログを書き込む
			logging.LogError("Your user type does not have permission to access this resource.", err)

			// エラーレスポンスを返す
			conversion.ResponseJson(ctx, http.StatusForbidden, nil)

			// 処理を終了する
			ctx.Abort()
			return
		}

		// mysqlのエラー
		logging.LogError("mysql error", err)

		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		// 処理を終了する
		ctx.Abort()
		return
	}

	// 次の処理へ
	ctx.Next()
}

// ユーザータイプが一般ユーザーであることを確認する
func AllowGeneralUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizeUserType(ctx, model.General)
	}
}

// ユーザータイプがレストランユーザーであることを確認する
func AllowRestaurantUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizeUserType(ctx, model.Restaurant)
	}
}
