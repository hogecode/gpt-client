package out

import (
	"fmt"
	"gpt-client/internal/application/port/out"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// QueryHistoryDB クエリ履歴のモデル
type QueryHistoryDB struct {
	ID        int    `gorm:"primary_key;auto_increment" json:"id"`
	Query     string `gorm:"type:text" json:"query"`
	Response  string `gorm:"type:text" json:"response"`
	CreatedAt string `json:"created_at"`
}

// HistoryRepository GORMを使った履歴の操作を行う構造体
type HistoryRepository struct {
	DB *gorm.DB
}

// NewHistoryRepository 新しい履歴リポジトリを作成
func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{DB: db}
}

// ListHistory DBの全てのクエリ履歴を返却
func (r *HistoryRepository) ListHistory() ([]out.QueryHistory, error) {
	var histories []QueryHistoryDB
	if err := r.DB.Find(&histories).Error; err != nil {
		return nil, fmt.Errorf("error retrieving history: %v", err)
	}

	// QueryHistoryからResponseフィールドに変換
	historiesWithResponse, err := mapQueryHistoryToResponse(histories)
	if err != nil {
		return nil, fmt.Errorf("error mapping query history: %v", err)
	}

	return historiesWithResponse, nil
}

// ListHistoryWithPagination ページネーションで履歴を返却
func (r *HistoryRepository) ListHistoryWithPagination(fromId int, toId int) ([]out.QueryHistory, error) {
	var histories []QueryHistoryDB
	if err := r.DB.Where("id BETWEEN ? AND ?", fromId, toId).Find(&histories).Error; err != nil {
		return nil, fmt.Errorf("error retrieving paginated history: %v", err)
	}

	// QueryHistoryからResponseフィールドに変換
	historiesWithResponse, err := mapQueryHistoryToResponse(histories)
	if err != nil {
		return nil, fmt.Errorf("error mapping query history: %v", err)
	}

	return historiesWithResponse, nil
}

// mapQueryHistoryToResponse DBのQueryHistoryDBから通常のQueryHistory構造体に変換
func mapQueryHistoryToResponse(histories []QueryHistoryDB) ([]out.QueryHistory, error) {
	var result []out.QueryHistory

	for _, history := range histories {
		// DBから取得したQueryHistoryDB構造体をQueryHistory構造体に変換
		result = append(result, out.QueryHistory{
			ID:       history.ID,
			Query:    history.Query,
		})
	}

	return result, nil
}

// InitializeDB 初期化とDB接続
func InitializeDB() (*gorm.DB, error) {
	// SQLiteのデータベースを開く
	db, err := gorm.Open("sqlite3", "./history.db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// テーブルのマイグレーション（存在しない場合は作成）
	if err := db.AutoMigrate(&QueryHistoryDB{}).Error; err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return db, nil
}
