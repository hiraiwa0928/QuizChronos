package main

import (
	"github.com/QuizChronos/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	routes.SetupQuizRoutes(r)

	// 8080ポートでサーバーを起動
	r.Run(":8080")
}
