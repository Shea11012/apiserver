package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("x-Request-Id")
		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}
		// 为了能在应用内使用，使用 set 暴露这个设置
		c.Set("x-Request-Id", requestId)
		// 设置响应头
		c.Writer.Header().Set("x-Request-Id", requestId)
	}
}
