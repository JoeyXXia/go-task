package response

import (
	"github.com/gin-gonic/gin"
)

// Success 返回一个成功的HTTP响应
func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"success": true,
		"data":    data,
	})
}

// Error 返回一个错误的HTTP响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"success": false,
		"error":   message,
	})
}
