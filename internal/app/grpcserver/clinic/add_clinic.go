package clinic

import (
	"context"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	localerrors "httpServer/internal/app/errors"
)

func (s *ServerClinic) AddClinic(
	ctx context.Context,
	req *pb.AddClinicRequest,
) (*pb.AddClinicResponse, error) {
	s.log.Info("AddClinic called")
	if errs := s.validator.AddClinic(req); errs != nil {
		err := localerrors.NewInvalidArgumentErr(*errs)
		s.log.InfoContext(ctx, "AddClinic validation error: ", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	useCaseReq := s.mapper.ProtoToAddClinicRequest(req)
	resp, err := s.clinicUseCase.AddClinic(ctx, useCaseReq)
	if err != nil {
		s.log.ErrorContext(ctx, "AddClinic UseCaseError", err.StatusCode().GRPC)
	}
	s.log.InfoContext(ctx, "Clinic added, clinic id: ", resp.ClinicId)
	response := s.mapper.AddClinicResponseToProtoResponse(*resp)
	return response, nil
}
