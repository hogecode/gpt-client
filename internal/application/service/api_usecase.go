package service

import (
	"context"
	"fmt"
	"gpt-client/internal/application/port/out"
)

type ApiUseCase struct {
	openAIApiGateway *out.OpenAIApiGateway
}

func NewApiUseCase(openAIApiGateway *out.OpenAIApiGateway) *ApiUseCase {
	return &ApiUseCase{openAIApiGateway: openAIApiGateway}
}

// SendQuery は、単一のクエリを送信し、結果を返します。
func (uc *ApiUseCase) SendQuery(ctx context.Context, query string) (out.OpenAIApiResponse, error) {
	// OpenAIApiGatewayのSendQueryメソッドを呼び出す
	response, err := uc.openAIApiGateway.SendQuery(ctx, query)
	if err != nil {
		return out.OpenAIApiResponse{}, fmt.Errorf("failed to send query: %w", err)
	}
	return response, nil
}

// SendQueriesFromFile は、ファイルから複数のクエリを読み込んで送信します。
func (uc *ApiUseCase) SendQueriesFromFile(ctx context.Context, filePath string) ([]out.OpenAIApiResponse, error) {
	// OpenAIApiGatewayのSendQueriesFromFileメソッドを呼び出す
	responses, err := uc.openAIApiGateway.SendQueriesFromFile(ctx, filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to send queries from file: %w", err)
	}
	return responses, nil
}

// PingApiKey は、APIキーが有効かどうかを確認します。
func (uc *ApiUseCase) PingApiKey() (bool, error) {
	// OpenAIApiGatewayのPingWithApiKeyメソッドを呼び出す
	valid, err := uc.openAIApiGateway.PingWithApiKey()
	if err != nil {
		return false, fmt.Errorf("failed to ping API key: %w", err)
	}
	return valid, nil
}