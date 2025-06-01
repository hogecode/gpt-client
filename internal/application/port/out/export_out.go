package out

type Format string

const (
    FormatJSON    Format = "json"
    FormatMarkdown Format = "markdown"
    FormatHTML     Format = "html"
)

type ExportUseCase interface {

	// DBの全てのクエリとその結果をFormat形式でエキスポート
    ExportHistory(format Format) error

	// DBの特定のクエリとその結果をFormat形式でエキスポート
    ExportHistoryWithPagination(fromId int, toId int, format Format) error
}