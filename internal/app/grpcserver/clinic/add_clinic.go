package clinic

import (
	"context"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerClinic) AddClinic(
	ctx context.Context,
	req *pb.AddClinicRequest,
) (*pb.AddClinicResponse, error) {
	s.log.Info("AddClinic called")
	if err := s.validator.AddClinic(req); err != nil {
		s.log.ErrorContext(ctx, "AddClinic validation erro", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	useCaseReq := s.mapper.ProtoToAddClinicRequest(req)
	resp, err := s.clinicUseCase.AddClinic(ctx, useCaseReq)
	if err != nil {
		s.log.ErrorContext(ctx, "AddClinic UseCaseError", err)
	}
	response := s.mapper.AddClinicResponseToProtoResponse(resp)
	return response, nil
}
