package routes

import "github.com/gin-gonic/gin"

func SetupQuizRoutes(r *gin.Engine) {
	r.GET("/quiz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"quiz-id": 1,
			"prob-id": 1,
		})
	})
}
