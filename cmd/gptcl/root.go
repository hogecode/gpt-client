package gptcl

import (
    "fmt"
    "github.com/spf13/cobra"
    "log"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "gptcl",
    Short: "gptcl is a command-line tool for interacting with GPT",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Use 'gptcl --help' to see available commands")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        log.Println(err)
        os.Exit(1)
    }
}

func init() {
    // サブコマンドの追加
    rootCmd.AddCommand(configCmd)
    rootCmd.AddCommand(queryCmd)
    rootCmd.AddCommand(historyCmd)
    rootCmd.AddCommand(exportCmd)
}

