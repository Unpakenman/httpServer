package grpcserver

import (
	"context"
	pb "github.com/Unpakenman/protos/gen/go/sso"
	rpc "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/grpcserver/clinic"
	"httpServer/internal/app/grpcserver/mapper"
	"httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
	"log/slog"
)

type ClinicServer struct {
	pb.UnimplementedClinicsServer
	inner *clinic.ServerClinic
}

func NewClinicServer(
	logger *slog.Logger,
	validator validator.Validator,
	mapper mapper.Mapper,
	clinicUseCase clinics.UseCase,
) pb.ClinicsServer {
	serverClinic := clinic.NewServer(logger, validator, mapper, clinicUseCase)
	return &ClinicServer{
		inner: serverClinic,
	}
}

func (s *ClinicServer) AddClinic(ctx context.Context, req *rpc.AddClinicRequest) (*rpc.AddClinicResponse, error) {
	return s.inner.AddClinic(ctx, req)
}

func (s *ClinicServer) AddEmployee(ctx context.Context, req *rpc.AddEmployeeRequest) (*rpc.AddEmployeeResponse, error) {
	return s.inner.AddEmployee(ctx, req)
}
