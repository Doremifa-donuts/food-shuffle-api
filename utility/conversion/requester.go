package conversion

import (
	"food-shuffle-api/utility/custom_error"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Content-TypeがJSONであることを確認し、値のバインドを行う
func BindJSON(ctx *gin.Context, obj any) error {
	// ヘッダーのContent-Typeにapplication/jsonが含まれているか確認
	if ctx.GetHeader("Content-Type") != "application/json" {
		// エラーレスポンスを返す
		return custom_error.NewError(http.StatusUnsupportedMediaType, "Content-Type is not application/json")
	}
	return ctx.ShouldBindJSON(&obj)
}
