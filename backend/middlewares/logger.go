package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエスト処理の前に実行される
		startTime := time.Now()

		// 次のミドルウェアまたはハンドラにリクエストを渡す
		c.Next()

		// リクエスト処理の後に実行される
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// ログを出力
		log.Printf("Request %s %s took %v", c.Request.Method, c.Request.URL, latency)
	}
}
