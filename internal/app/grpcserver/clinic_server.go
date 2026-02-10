package grpcserver

import (
	pb "github.com/Unpakenman/protos/gen/go/sso"
	"httpServer/internal/app/grpcserver/mapper"
	"httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
	"log/slog"
)

type ClinicServer struct {
	pb.UnimplementedClinicsServer
	log       *slog.Logger
	validator validator.Validator
	mapper    mapper.Mapper
	usecase   clinics.UseCase
}

func NewClinicServer(
	logger *slog.Logger,
	validator validator.Validator,
	mapper mapper.Mapper,
	clinicUseCase clinics.UseCase,
) pb.ClinicsServer {
	return &ClinicServer{
		log:       logger,
		validator: validator,
		mapper:    mapper,
		usecase:   clinicUseCase,
	}
}
