package main

import (
	"gpt-client/cmd/gptcl"
	"gpt-client/internal/adapters/out"
	"log"
)

func main() {
	// DBの初期化
	db, err := out.InitializeDB()
	if err != nil {
		log.Fatalf("Error initializing DB: %v", err)
	}
	defer db.Close()

	// 依存関係の設定
	// adapters/outフォルダの初期化
	historyRepository := out.NewHistoryRepository(db)

	// サービスの初期化
	// サービスに上記の依存関係を注入
	// historyService := service.NewHistoryService(historyRepository)

	// コマンドにサービスを注入
	// gptcl.RegisterConfigCommands(historyService)

	// CLIの実行
	gptcl.Execute()
}
