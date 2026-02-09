package grpcserver

import (
	pb "github.com/Unpakenman/protos/gen/go/sso"
	"httpServer/internal/app/grpcserver/mapper"
	"httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
)

type ClinicServer struct {
	pb.UnimplementedClinicsServer
	validator validator.Validator
	mapper    mapper.Mapper
	usecase   clinics.UseCase
}

func NewClinicServer(
	validator validator.Validator,
	mapper mapper.Mapper,
	clinicUseCase clinics.UseCase,
) pb.ClinicsServer {
	return &ClinicServer{
		validator: validator,
		mapper:    mapper,
		usecase:   clinicUseCase,
	}
}
