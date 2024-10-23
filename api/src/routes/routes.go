package routes

import (
	"github.com/QuizChronos/service"
	"github.com/gin-gonic/gin"
)

// SetupRoutes はアプリケーションのルートを設定します
func SetupRoutes(r *gin.Engine) {
	r.GET("/quiz", service.GetQuiz)
}
