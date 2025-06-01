package out

type QueryHistory struct {
    ID      int
    Query   string
}

type HistoryService interface {

	// DBの全てのクエリを返却
    ListHistory() ([]QueryHistory, error)

	// DBの特定のクエリを返却
    ListHistoryWithPagination(fromId int, toId int) ([]QueryHistory, error)
}
