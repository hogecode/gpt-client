package gptcl

import (
    "fmt"
    "github.com/spf13/cobra"
    "gpt-client/internal/application"
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

func init() {
    configCmd.Flags().StringVar(&apiKey, "api-key", "", "Set the API key")
    configCmd.Flags().StringVar(&outputDir, "output-dir", "", "Set the output directory")
}
