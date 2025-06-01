package out

import (
	"context"
	"fmt"
	"gpt-client/internal/application/port/out"
	"io/ioutil"

	"github.com/go-resty/resty/v2"
)

const (
	baseURL = "https://api.openai.com/v1"
)

// OpenAIApiGatewayImpl 構造体の定義
type OpenAIApiGatewayImpl struct {
	client  *resty.Client
	apiKey  string
	baseURL string
}

// NewOpenAIApiGatewayImpl コンストラクタ関数
func NewOpenAIApiGatewayImpl(apiKey string) *OpenAIApiGatewayImpl {
	client := resty.New()

	// APIキーを設定
	client.SetHeader("Authorization", "Bearer "+apiKey)

	return &OpenAIApiGatewayImpl{
		client:  client,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// SendQuery 単一のクエリを送信するメソッド
func (c *OpenAIApiGatewayImpl) SendQuery(ctx context.Context, query string) (out.OpenAIApiResponse, error) {
	// リクエストの作成
	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(map[string]string{"prompt": query, "max_tokens": "150"}).
		SetResult(&out.OpenAIApiResponse{}).
		Post(c.baseURL + "/completions")

	if err != nil {
		return out.OpenAIApiResponse{}, fmt.Errorf("failed to send query: %w", err)
	}

	// レスポンスの処理
	apiResponse := resp.Result().(*out.OpenAIApiResponse)
	return *apiResponse, nil
}

// SendQueriesFromFile ファイルから複数のクエリを読み込み、送信するメソッド
func (c *OpenAIApiGatewayImpl) SendQueriesFromFile(ctx context.Context, filePath string) ([]out.OpenAIApiResponse, error) {
	// ファイルの読み込み
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// ファイル内の各行をクエリとして送信
	var responses []out.OpenAIApiResponse
	queries := string(fileData)
	for _, query := range queries {
		response, err := c.SendQuery(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("failed to send query from file: %w", err)
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// PingWithApiKey APIキーが有効か確認するメソッド
func (c *OpenAIApiGatewayImpl) PingWithApiKey() (bool, error) {
	// ヘルスチェックとして、適切なエンドポイントにリクエストを送信
	resp, err := c.client.R().
		SetHeader("Authorization", "Bearer "+c.apiKey).
		Get(c.baseURL + "/ping")

	if err != nil {
		return false, fmt.Errorf("failed to ping with API key: %w", err)
	}

	// APIから成功レスポンスが返ってきた場合
	if resp.StatusCode() == 200 {
		return true, nil
	}

	return false, fmt.Errorf("API key is invalid or the server is unreachable")
}
