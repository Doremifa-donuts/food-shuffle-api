package conversion

import (
	"food-shuffle-api/utility/custom_error"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Content-TypeがJSONであることを確認し、値のバインドを行う
func BindJSON(ctx *gin.Context, objects ...any) *custom_error.CustomError {
	// ヘッダーのContent-Typeにapplication/jsonが含まれているか確認
	if ctx.GetHeader("Content-Type") != "application/json" {
		// エラーレスポンスを返す
		return custom_error.NewError(http.StatusUnsupportedMediaType, "Content-Type is not application/json")
	}

	// 構造体に値をバインドする
	for _, obj := range objects {
		err := ctx.ShouldBindBodyWithJSON(&obj)
		if err != nil {
			return custom_error.NewError(http.StatusBadRequest, "failed bin to json")
		}
	}
	return nil
}

// 
