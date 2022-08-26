package sql

import (
	"github.com/gin-gonic/gin"
)

var auth gin.HandlerFunc = func(context *gin.Context) {
	value, err := context.Cookie("gin_cookie")
	if err != nil {
		context.JSON(403, gin.H{
			"message": "认证失败,没有cookie",
		})
		context.Abort()
	} else {
		context.Set("cookie", value)
		context.Next()
	}
}
