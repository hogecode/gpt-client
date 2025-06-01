package gptcl

import (
    "fmt"
    "github.com/spf13/cobra"
    "gpt-client/internal/application"
)

var queryText string
var format string

var queryCmd = &cobra.Command{
    Use:   "query",
    Short: "Send a query to GPT",
    Run: func(cmd *cobra.Command, args []string) {
        // クエリを送信
        err := application.SendQuery(queryText, format)
        if err != nil {
            fmt.Println("Error sending query:", err)
            return
        }
        fmt.Println("Query sent successfully.")
    },
}

func init() {
    queryCmd.Flags().StringVar(&queryText, "query", "", "Query text")
    queryCmd.Flags().StringVar(&format, "format", "markdown", "Output format")
}
