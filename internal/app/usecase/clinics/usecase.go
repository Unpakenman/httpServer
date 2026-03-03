package clinics

import (
	"context"
	"httpServer/internal/app/config"
	localerrors "httpServer/internal/app/errors"
	ihttpservice "httpServer/internal/app/internal_services/internal_http_service"
	logger "httpServer/internal/app/log"
	"httpServer/internal/app/provider"
	providerCache "httpServer/internal/app/provider/cache"
	"httpServer/internal/app/rabbitmq_service"
)

type clinicsUseCase struct {
	provider        provider.GoExampleProvider
	logger          logger.LogClient
	internalService ihttpservice.Service
	rmqService      rabbitmq_service.RMQService
	config          *config.Values
	cache           providerCache.Cache
}

func NewUseCase(
	provider provider.GoExampleProvider,
	logger logger.LogClient,
	internalService ihttpservice.Service,
	rmqService rabbitmq_service.RMQService,
	config *config.Values,
	providerCache providerCache.Cache,
) UseCase {
	return &clinicsUseCase{
		provider:        provider,
		logger:          logger,
		internalService: internalService,
		rmqService:      rmqService,
		config:          config,
		cache:           providerCache,
	}
}

type UseCase interface {
	CreatePatient(ctx context.Context, req CreatePatientRequest) (CreatePatientResponse, localerrors.Error)
	AddClinic(ctx context.Context, req AddClinicRequest) (*AddClinicResponse, localerrors.Error)
	AddEmployee(ctx context.Context, req AddEmployeeRequest) (*AddEmployeeResponse, localerrors.Error)
	AddAppointment(ctx context.Context, req AddAppointmentRequest) (*AddAppointmentResponse, localerrors.Error)
	CreateNewService(ctx context.Context, request CreateNewServiceRequest) localerrors.Error
	AppointmentList(ctx context.Context) localerrors.Error
	AppointmentsSlots(ctx context.Context, req AppointmentsSlotsRequest) (*AppointmentsSlotsResponse, localerrors.Error)
}
