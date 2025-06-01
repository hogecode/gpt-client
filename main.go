package main

import (
	"gpt-client/cmd/gptcl"
	"gpt-client/internal/adapters/out"
	"gpt-client/internal/application/service"
	"log"
)

func main() {
	// DBの初期化
	db, err := out.InitializeDB()
	if err != nil {
		log.Fatalf("Error initializing DB: %v", err)
	}
	defer db.Close()

	var mockKey string

	// 依存関係の設定
	// adapters/outフォルダの初期化
	historyRepository := out.NewHistoryRepositoryImpl(db)
	apiClient := out.NewOpenAIApiGatewayImpl(mockKey)

	// application/serviceフォルダの初期化
	// サービスに上記の依存関係を注入
	apiService := service.NewApiUseCase(apiClient)

	// cmd/gptclフォルダにサービスを注入
	// gptcl.RegisterConfigCommands(historyService)

	// CLIの実行
	gptcl.Execute()
}
