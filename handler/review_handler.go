package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"

	"food-shuffle-api/model"
	"food-shuffle-api/service"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"
)

// サービス層のメソッドは構造体と紐づいて管理されているため、処理を投げる構造体を呼び出す
var ReviewService = service.ReviewService{}

// レビューを投稿する
func ReviewPostHandler(ctx *gin.Context) {

	// リクエストを構造体にバインドする
	var review model.Review
	// idを構造体にバインドする
	uuid, ok := ctx.Get("uuid")
	if !ok {
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	// uuidを構造体にバインドする
	review.UserUuid = uuid.(string)

	// リクエストを構造体にバインドする
	err := conversion.RequestSaveImagesAndBindJSON(ctx, &review)
	if err != nil {
		// カスタムエラーの変数を宣言する
		var customError *custom_error.CustomError

		// カスタムエラーかどうかを確認する
		if errors.As(err, &customError) {
			// エラーを返す
			conversion.ResponseJson(ctx, customError.StatusCode(), nil)
			return
		} else { // TODO: カスタムエラー以外のエラーを分類する
			// エラーを返す
			conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
			return
		}
	}

	// サービス層のメソッドを呼び出す
	reviewUuid, err := ReviewService.PostReview(review)
	if err != nil {

		// カスタムエラーの変数を宣言する
		var customError *custom_error.CustomError

		// カスタムエラーかどうかを確認する
		if errors.As(err, &customError) { // カスタムエラー
			conversion.ResponseJson(ctx, customError.StatusCode(), nil)
			return
		}

		// mysqlのエラーを分類する
		var mysqlError *mysql.MySQLError

		if errors.As(err, &mysqlError) {
			switch mysqlError.Number {
			case 1452: // 外部キー制約 	レストランIDまたはユーザーIDが不正である場合に発生する
				conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
				return
			default:
				conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
				return
			}
		}

		// それ以外のエラー
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// レスポンスを返す
	conversion.ResponseJson(ctx, http.StatusOK, reviewUuid)

}
