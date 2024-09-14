package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GreenTeaProgrammers/MatsuriSNS/ent"

	"github.com/GreenTeaProgrammers/MatsuriSNS/config"
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

	// ユーザーを作成できるか確認用
	user, err := client.User.
		Create().
		SetUsername("testuser").
		SetEmail("testuser2@example.com").
		SetPasswordHash("hashed_password").
		Save(context.Background())
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	fmt.Printf("User created: %v\n", user)
}
