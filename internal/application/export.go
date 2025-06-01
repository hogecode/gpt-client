package application

import (
    "fmt"
    "os"
)

func ExportHistoryToFile(fileName string) error {
    historyRecords, err := GetAllHistory()
    if err != nil {
        return fmt.Errorf("could not fetch history: %v", err)
    }

    file, err := os.Create(fileName)
    if err != nil {
        return fmt.Errorf("could not create file: %v", err)
    }
    defer file.Close()

    for _, record := range historyRecords {
        _, err := fmt.Fprintf(file, "Query: %s\nResponse: %s\n\n", record.Query, record.Response)
        if err != nil {
            return fmt.Errorf("could not write to file: %v", err)
        }
    }

    return nil
}
