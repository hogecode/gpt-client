package gptcl

import (
    "fmt"
    "github.com/spf13/cobra"
    "gpt-client/internal/application"
)

var historyCmd = &cobra.Command{
    Use:   "history",
    Short: "Show interaction history",
    Run: func(cmd *cobra.Command, args []string) {
        // 履歴を表示
        historyRecords, err := application.GetAllHistory()
        if err != nil {
            fmt.Println("Error fetching history:", err)
            return
        }

        for _, record := range historyRecords {
            fmt.Println(record)
        }
    },
}
