package internal_http_service

import (
	"context"
	httpclient "httpServer/internal/app/client/http"
	"httpServer/internal/app/config"
)

type internalServiceAPI struct {
	baseURL    string
	httpClient httpclient.Client
}

//go:generate ../../../../bin/mockery --with-expecter --case=underscore --name=Service

type Service interface {
	SomeExample(ctx context.Context, req ActionRequest) (*ActionResponse, error)
}

func NewService(cfg *config.SomeHttpServiceConfig, client httpclient.Client) Service {

	return &internalServiceAPI{
		baseURL:    cfg.URL,
		httpClient: client,
	}
}
