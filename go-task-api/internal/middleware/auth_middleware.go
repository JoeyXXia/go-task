package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joey/go-task/pkg/jwt"
	"github.com/joey/go-task/pkg/response"
)

// AuthMiddleware 验证请求是否包含有效的JWT令牌
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取令牌
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// 如果请求头中没有令牌，则尝试从Cookie中获取
			tokenCookie, err := c.Cookie("token")
			if err != nil {
				response.Error(c, http.StatusUnauthorized, "Unauthorized: no authentication token provided")
				c.Abort()
				return
			}
			authHeader = "Bearer " + tokenCookie
		}

		// 检查令牌格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "Unauthorized: invalid token format")
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 验证令牌
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Unauthorized: "+err.Error())
			c.Abort()
			return
		}

		// 将用户ID存储在上下文中，以便后续处理器使用
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
