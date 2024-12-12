package middleware

import (
	"mini-wallet-api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "fail",
				"data": gin.H{
					"error": "Authorization header missing or invalid",
				},
			})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "fail",
				"data": gin.H{
					"error": "Invalid or expired token",
				},
			})
			c.Abort()
			return
		}

		c.Set("customer_xid", claims.CustomerXID)
		c.Next()
	}
}
