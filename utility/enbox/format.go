package enbox

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseJson(c *gin.Context, httpStatus int, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"response": gin.H{
			"status":  http.StatusText(httpStatus),
			"data":    data,
		},
	})
}