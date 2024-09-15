package main

import (
	"context"
	"log"
	"net/http"

	"github.com/GreenTeaProgrammers/MatsuriSNS/config"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent"
	"github.com/GreenTeaProgrammers/MatsuriSNS/middlewares"
	"github.com/GreenTeaProgrammers/MatsuriSNS/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 環境変数を読み込む
	config.LoadEnv()

	// DSNを取得してデータベース接続を設定
	dsn := config.GetDSN()

	// Entクライアントを作成
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open connection to database: %v", err)
	}
	defer client.Close()

	// データベーススキーマをマイグレーション
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	// Ginのインスタンスを作成
	r := gin.Default()

	// ミドルウェアの適用
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.RecoveryMiddleware())

	// 確認用のルート
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from MatsuriSNS"})
	})

	// 認証ルートの設定
	routes.SetupAuthRoutes(r, client)
	routes.SetupPostRoutes(r, client)

	// サーバーをポート8080で起動
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run Gin server: %v", err)
	}
}
