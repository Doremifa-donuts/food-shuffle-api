package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"

	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/service"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"
)

// サービス層のメソッドは構造体と紐づいて管理されているため、処理を投げる構造体を呼び出す
var ReviewService = service.ReviewService{}

// すれ違いで受け取ったレビューの一覧を取得する
func GetReceivedReviewsByUserHandler(ctx *gin.Context) {
	// ユーザーIDを取得する
	uuid, _ := ctx.Get("uuid")
	// 型変換
	uuidAdjusted := uuid.(string)

	// レビュー一覧を取得するサービスに投げる
	reviews, err := ReviewService.GetNewReviewsByUser(uuidAdjusted)
	// エラーハンドリング
	if err != nil {
		// エラーを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// レビュー一覧を返す
	conversion.ResponseJson(ctx, http.StatusOK, reviews)
}

// 　興味ありに保存されたレビューの一覧を取得する
func GetInterestedReviewsByUserHandler(ctx *gin.Context) {
	// ユーザーIDを取得する
	uuid, _ := ctx.Get("uuid")
	// 型変換
	uuidAdjusted := uuid.(string)

	// レビュー一覧を取得するサービスに投げる
	reviews, err := ReviewService.GetInterestedReviewsByUser(uuidAdjusted)
	// エラーハンドリング
	if err != nil {
		// エラーを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// レビュー一覧を返す
	conversion.ResponseJson(ctx, http.StatusOK, reviews)
}

// いいねをしたレビューの一覧を取得する
func GetLikedReviewsByUserHandler(ctx *gin.Context) {
	// ユーザーIDを取得する
	uuid, _ := ctx.Get("uuid")
	// 型変換
	uuidAdjusted := uuid.(string)

	// レビュー一覧を取得するサービスに投げる
	reviews, err := ReviewService.GetLikedReviewsByUser(uuidAdjusted)
	// エラーハンドリング
	if err != nil {
		// エラーを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// レビュー一覧を返す
	conversion.ResponseJson(ctx, http.StatusOK, reviews)
}

// レビューのステータスを更新する
func PutReviewStatusByUserHandler(ctx *gin.Context) {
	// リクエストを構造体にバインドする
	var userReviewFlag model.UserReviewFlag

	// リクエストを構造体にバインドする
	userUuid, _ := ctx.Get("uuid")
	userReviewFlag.UserUuid = userUuid.(string)
	fmt.Println("userReviewFlag.UserUuid:", userReviewFlag.UserUuid)

	// レビューUUIDが存在しているかを確認する
	reqReviewUuid := ctx.Param("review_uuid")
	if reqReviewUuid == "" {
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}
	// レビューUUIDを構造体に格納する
	userReviewFlag.ReviewUuid = reqReviewUuid

	fmt.Println("userReviewFlag.ReviewUuid:", userReviewFlag.ReviewUuid)

	// リクエストステータスが適切かどうかを確認する
	reqStatus := ctx.Param("review_status")
	switch reqStatus {
	case "interested":
		userReviewFlag.ReviewStatus = model.Interested
	case "not_interested":
		userReviewFlag.ReviewStatus = model.NotInterested
	case "liked":
		userReviewFlag.ReviewStatus = model.Iiked
	default:
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	// レビューのステータスを更新するサービスに投げる
	err := ReviewService.UpdateReviewStatus(userReviewFlag)
	// エラーハンドリング
	if err != nil {
		// カスタムエラーの変数を宣言する
		var customError *custom_error.CustomError

		// カスタムエラーかどうかを確認する
		if errors.As(err, &customError) {
			// カスタムエラーを返す
			conversion.ResponseJson(ctx, customError.StatusCode(), nil)
			return
		}
		// 分類していないエラーを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}
	conversion.ResponseJson(ctx, http.StatusOK, nil)
}

// レビューを投稿する
func PostReviewByUserHandler(ctx *gin.Context) {

	// multipart/form-dataであることを確認する
	if ctx.ContentType() != "multipart/form-data" {
		logging.LogError("invalid content type", nil)
		conversion.ResponseJson(ctx, http.StatusUnsupportedMediaType, nil)
	}

	// リクエストを構造体にバインドする
	var review model.Review
	// idを構造体にバインドする
	uuid, _ := ctx.Get("uuid")

	// uuidを構造体にバインドする
	review.UserUuid = uuid.(string)

	// リクエストを受け取る
	form, err := ctx.MultipartForm()
	if err != nil {
		logging.LogError("failed get multi part form", err)
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
	}

	// jsonを構造体にバインド
	// jsonデータを取得
	jsonData := form.Value["data"][0]
	// jsonデータが空の場合はエラーを返す
	if jsonData == "" {
		logging.LogError("Invalid JSON data:", nil)
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	} else {
		// JSONデータを構造体にパース
		if err := json.Unmarshal([]byte(jsonData), &review); err != nil {
			logging.LogError("Error parsing JSON data:", err)
			conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
			return
		}
	}

	// 画像をスライスに格納
	images := form.File["images[]"]

	if len(images) > 10 && len(images) <= 0 {
		// 画像なし
		logging.LogError("images[] length is too long or zero", nil)
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	// サービス層のメソッドを呼び出す
	reviewUuid, err := ReviewService.PostReview(review, images)
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

// ユーザーがシェアするレビューを設定する
func PutReviewShareSettingHandler(ctx *gin.Context) {

	// ヘッダーのContent-Typeにapplication/jsonが含まれているか確認
	if ctx.GetHeader("Content-Type") != "application/json" {
		logging.LogError("Content-Type is not application/json", nil)

		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusUnsupportedMediaType, nil)
		return
	}

	// ユーザーUUIDを取得
	userUuid, _ := ctx.Get("uuid")

	// 構造体にバインド
	var bShareSettingReview model.ShareSettingReview
	err := ctx.ShouldBindJSON(&bShareSettingReview)
	if err != nil {
		logging.LogError("could not bind to json", nil)
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	bShareSettingReview.UserUuid = userUuid.(string)

	// サービスに処理を投げる
	res, err := ReviewService.SetShareReview(bShareSettingReview)
	if err != nil {
		// カスタムエラーの場合のレスポンス
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}
		// その他の場合のレスポンス
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
	}
	conversion.ResponseJson(ctx, http.StatusOK, res)

}

func GetPostedReviewHandler(ctx *gin.Context) {
	restaurantUuid := ctx.Param("restaurant_uuid")
	if restaurantUuid == "" {
		logging.LogError("RestaurantUuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	userUuid, _ := ctx.Get("uuid")
	if userUuid == nil {
		logging.LogError("UserUuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	// サービスに処理を投げる
	res, err := ReviewService.GetReviewDetail(restaurantUuid, userUuid.(string))
	if err != nil {
		logging.LogError("get review detail failed", err)

		// カスタムエラーを分類する
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}

		// その他の場合のエラーレスポンス
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// 成功レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, res)
}
