package service

import "github.com/gin-gonic/gin"

func GetQuiz(c *gin.Context) {
	c.JSON(200, gin.H{
		"quiz-id": 1,
		"prob-id": 1,
	})
}
