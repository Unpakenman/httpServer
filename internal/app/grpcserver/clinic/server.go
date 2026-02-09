package clinic

import (
	"httpServer/internal/app/grpcserver/mapper"
	"httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
)

type ServerClinic struct {
	validator     validator.Validator
	mapper        mapper.Mapper
	clinicUseCase clinics.UseCase
}

func NewServer(
	validator validator.Validator,
	mapper mapper.Mapper,
	clinicUseCase clinics.UseCase,
) *ServerClinic {
	return &ServerClinic{
		validator:     validator,
		mapper:        mapper,
		clinicUseCase: clinicUseCase,
	}
}
