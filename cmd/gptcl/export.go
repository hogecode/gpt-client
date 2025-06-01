package gptcl

import (
    "fmt"
    "github.com/spf13/cobra"
    "gpt-client/internal/application"
)

var exportCmd = &cobra.Command{
    Use:   "export",
    Short: "Export history to a file",
    Run: func(cmd *cobra.Command, args []string) {
        // エクスポート処理
        err := application.ExportHistoryToFile("history.md")
        if err != nil {
            fmt.Println("Error exporting history:", err)
            return
        }
        fmt.Println("History exported successfully.")
    },
}
