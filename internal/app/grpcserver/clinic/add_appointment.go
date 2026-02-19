package clinic

import (
	"context"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	localerrors "httpServer/internal/app/errors"
)

func (s *ServerClinic) AddAppointment(
	ctx context.Context,
	req *pb.AddAppointmentRequest,
) (*pb.AddAppointmentResponse, error) {
	s.log.InfoCtx(ctx, "AddAppointment called")
	if errs := s.validator.AddAppointment(req); errs != nil {
		err := localerrors.NewInvalidArgumentErr(*errs)
		s.log.InfoCtx(ctx, "AddClinic validation error: ", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	useCaseReq, err := s.mapper.ProtoToAddAppointmentRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	useCaseResp, err := s.clinicUseCase.AddAppointment(ctx, useCaseReq)
	if err != nil {
		s.log.ErrorCtx(ctx, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	response := s.mapper.AddAppointmentResponseToProtoResponse(useCaseResp)
	return response, nil
}
