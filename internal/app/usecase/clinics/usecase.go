package clinics

import (
	"context"
	ihttpservice "httpServer/internal/app/internal_services/internal_http_service"
	"httpServer/internal/app/provider"
)

type clinicsUseCase struct {
	provider        provider.GoExampleProvider
	internalService ihttpservice.Service
}

func NewUseCase(
	provider provider.GoExampleProvider,
	internalService ihttpservice.Service,
) UseCase {
	return &clinicsUseCase{
		provider:        provider,
		internalService: internalService,
	}
}

type UseCase interface {
	CreatePatient(
		ctx context.Context,
		req CreatePatientRequest) (CreatePatientResponse, error)
}
