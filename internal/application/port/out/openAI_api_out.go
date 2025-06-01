package out

import "context"

type OpenAIApiGateway interface {
	// SendQuery は、単一のクエリを送信
	SendQuery(ctx context.Context, query string) (OpenAIApiResponse, error)

	// ファイル内の複数のクエリを送信
	SendQueriesFromFile(ctx context.Context, filePath string) ([]OpenAIApiResponse, error)

    // APIキーが有効か確かめる
    PingWithApiKey() (bool, error)
}

type OpenAIApiResponse struct {
    Text    string `json:"text"`    
    Error   string `json:"error"`   
    Success bool   `json:"success"` 
}
