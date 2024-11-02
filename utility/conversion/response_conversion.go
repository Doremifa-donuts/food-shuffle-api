package conversion

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
// データを適切な形式に変換する処理をまとめて記述する

// レスポンスは全てJSON形式の入れ子構造にする
func ResponseJson(ctx *gin.Context, httpStatus int, data interface{}) {
	ctx.JSON(httpStatus, gin.H{
		"Response": gin.H{
			"Status":  http.StatusText(httpStatus),
			"Data":    data,
		},
	})
}