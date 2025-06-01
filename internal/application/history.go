package application

import (
    // 実際にはDB操作に必要なパッケージをインポート
)

type HistoryRecord struct {
    ID      int
    Query   string
    Response string
}

func GetAllHistory() ([]HistoryRecord, error) {
    // ここでSQLiteから履歴を取得
    // 実際にはGORMやSQLクエリを使う

    return []HistoryRecord{
        {ID: 1, Query: "こんにちは", Response: "こんにちは！"},
        {ID: 2, Query: "お元気ですか？", Response: "元気です、ありがとう！"},
    }, nil
}
