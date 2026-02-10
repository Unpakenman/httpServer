package clinics

import (
	"context"
	"httpServer/internal/app/config"
	ihttpservice "httpServer/internal/app/internal_services/internal_http_service"
	"httpServer/internal/app/provider"
	"log/slog"
)

type clinicsUseCase struct {
	provider        provider.GoExampleProvider
	logger          slog.Logger
	internalService ihttpservice.Service
	config          *config.Values
}

func NewUseCase(
	provider provider.GoExampleProvider,
	logger slog.Logger,
	internalService ihttpservice.Service,
	config *config.Values,
) UseCase {
	return &clinicsUseCase{
		provider:        provider,
		logger:          logger,
		internalService: internalService,
		config:          config,
	}
}

type UseCase interface {
	CreatePatient(
		ctx context.Context,
		req CreatePatientRequest) (CreatePatientResponse, error)
	AddClinic(
		ctx context.Context,
		req AddClinicRequest) (AddClinicResponse, error)
}
