package clinics

import (
	"context"
	"httpServer/internal/app/config"
	localerrors "httpServer/internal/app/errors"
	ihttpservice "httpServer/internal/app/internal_services/internal_http_service"
	logger "httpServer/internal/app/log"
	"httpServer/internal/app/provider"
)

type clinicsUseCase struct {
	provider        provider.GoExampleProvider
	logger          logger.LogClient
	internalService ihttpservice.Service
	config          *config.Values
}

func NewUseCase(
	provider provider.GoExampleProvider,
	logger logger.LogClient,
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
	CreatePatient(ctx context.Context, req CreatePatientRequest) (CreatePatientResponse, localerrors.Error)
	AddClinic(ctx context.Context, req AddClinicRequest) (*AddClinicResponse, localerrors.Error)
	AddEmployee(ctx context.Context, req AddEmployeeRequest) (*AddEmployeeResponse, localerrors.Error)
	AddAppointment(ctx context.Context, req AddAppointmentRequest) (*AddAppointmentResponse, localerrors.Error)
}
