package middleware

import "github.com/gin-gonic/gin"

func Sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			panic("token不能为空")
		}

		c.Next()
	}
}
