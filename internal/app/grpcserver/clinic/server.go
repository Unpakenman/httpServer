package clinic

import (
	"httpServer/internal/app/grpcserver/mapper"
	logger "httpServer/internal/app/log"
	"httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
)

type ServerClinic struct {
	log           logger.LogClient
	validator     validator.Validator
	mapper        mapper.Mapper
	clinicUseCase clinics.UseCase
}

func NewServer(
	logger logger.LogClient,
	validator validator.Validator,
	mapper mapper.Mapper,
	clinicUseCase clinics.UseCase,
) *ServerClinic {
	return &ServerClinic{
		log:           logger,
		validator:     validator,
		mapper:        mapper,
		clinicUseCase: clinicUseCase,
	}
}
