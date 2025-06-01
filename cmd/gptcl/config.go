package gptcl

import (
	"fmt"
	"gpt-client/internal/application"
	"gpt-client/internal/application/service"

	"github.com/spf13/cobra"
)

var apiKey string
var outputDir string

var configCmd = &cobra.Command{
    Use:   "config",
    Short: "Manage configuration settings",
    Run: func(cmd *cobra.Command, args []string) {
        // APIキーを保存する処理
        if apiKey != "" {
            err := application.SaveApiKey(apiKey)
            if err != nil {
                fmt.Println("Error saving API Key:", err)
                return
            }
            fmt.Println("API Key saved successfully.")
        }

		// 保存先のフォルダの設定を保存する処理
        if outputDir != "" {
            err := application.SaveOutputDir(outputDir)
            if err != nil {
                fmt.Println("Error saving output directory:", err)
                return
            }
            fmt.Println("Output directory saved successfully.")
        }
    },
}

// 依存関係をコマンドに渡すための関数
// TODO: port/inのインターフェースに変えることを考える
func RegisterConfigCommands(configService service.ConfigUseCase) {
	// configCmdに必要なサービスを渡す
	configCmd.Flags().StringVar(&apiKey, "api-key", "", "Set the API key")
	configCmd.Flags().StringVar(&outputDir, "output-dir", "", "Set the output directory")
}
