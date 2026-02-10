package clinic

import (
	"httpServer/internal/app/grpcserver/mapper"
	"httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
	"log/slog"
)

type ServerClinic struct {
	log           *slog.Logger
	validator     validator.Validator
	mapper        mapper.Mapper
	clinicUseCase clinics.UseCase
}

func NewServer(
	logger *slog.Logger,
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
