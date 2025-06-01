package out

import (
	"encoding/json"
	"fmt"
	"gpt-client/internal/application/port/out"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ConfigServiceImpl 構造体は ConfigService インターフェースを実装
type ConfigServiceImpl struct {
	configFilePath string
}

// 新しい ConfigServiceImpl を作成する関数
func NewConfigServiceImpl() *ConfigServiceImpl {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
	}

	// 設定ファイルのパスを決定
	configFilePath := filepath.Join(homeDir, ".gptcl", "config.json")

	return &ConfigServiceImpl{
		configFilePath: configFilePath,
	}
}

// JSONファイルから設定を読み込む
func (c *ConfigServiceImpl) loadConfig() (*out.Configuration, error) {
	// 設定ファイルが存在しない場合は新しい設定を返す
	if _, err := os.Stat(c.configFilePath); os.IsNotExist(err) {
		return &out.Configuration{}, nil
	}

	// 設定ファイルを読み込む
	data, err := ioutil.ReadFile(c.configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	// 設定を解析して返す
	var config out.Configuration
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &config, nil
}

// 設定をJSONファイルに保存する
func (c *ConfigServiceImpl) saveConfig(config *out.Configuration) error {
	// ディレクトリが存在しない場合、作成する
	dir := filepath.Dir(c.configFilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
		// 設定をJSONとして保存
	}

	file, err := os.Create(c.configFilePath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(config)
}

// SetApiKey はAPIキーを設定し、JSONに保存する
func (c *ConfigServiceImpl) SetApiKey(apiKey string) error {
	config, err := c.loadConfig()
	if err != nil {
		return err
	}

	// APIキーを設定
	config.ApiKey = apiKey

	// 設定を保存
	return c.saveConfig(config)
}

// SetOutputDir は出力ディレクトリを設定し、JSONに保存する
func (c *ConfigServiceImpl) SetOutputDir(outputDir string) error {
	config, err := c.loadConfig()
	if err != nil {
		return err
	}

	// 出力ディレクトリを設定
	config.OutputDir = outputDir

	// 設定を保存
	return c.saveConfig(config)
}

// ShowConfig は現在の設定を返す
func (c *ConfigServiceImpl) ShowConfig() (*out.Configuration, error) {
	return c.loadConfig()
}