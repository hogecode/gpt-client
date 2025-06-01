package out

type Configuration struct {
    ApiKey    string `json:"api_key"`
    OutputDir string `json:"output_dir"`
}

type ConfigService interface {
	
	// JSONにAPIキーを保存
    SetApiKey(apiKey string) error

	// JSONにエキスポートフォルダを保存
    SetOutputDir(outputDir string) error

	// JSONの内容を取得
    ShowConfig() (*Configuration, error)
}
