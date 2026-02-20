package grpcserver

import (
	pb "github.com/Unpakenman/protos/gen/go/sso"
	"httpServer/internal/app/grpcserver/clinic"
	"httpServer/internal/app/grpcserver/mapper"
	logger "httpServer/internal/app/log"
	"httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
)

type ClinicServer struct {
	*clinic.ServerClinic
}

func NewClinicServer(
	logger logger.LogClient,
	validator validator.Validator,
	mapper mapper.Mapper,
	clinicUseCase clinics.UseCase,
) pb.ClinicsServer {
	return &ClinicServer{
		clinic.NewServer(logger, validator, mapper, clinicUseCase),
	}
}
