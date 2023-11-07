package api_server

import (
	"github.com/disism/saikan/internal/authx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := authx.ValidateJWT(c); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": err.Error(),
			})
			return
		}

		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
