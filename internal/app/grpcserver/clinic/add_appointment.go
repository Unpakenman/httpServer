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
	s.log.InfoContext(ctx, "AddAppointment called")
	if errs := s.validator.AddAppointment(req); errs != nil {
		err := localerrors.NewInvalidArgumentErr(*errs)
		s.log.InfoContext(ctx, "AddClinic validation error: ", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	useCaseReq := s.mapper.ProtoToAddAppointmentRequest(req)
	useCaseResp, err := s.clinicUseCase.AddAppointment(ctx, useCaseReq)
	if err != nil {
		s.log.ErrorContext(ctx, "AddClinic AddAppointment error: ", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	response := s.mapper.AddAppointmentResponseToProtoResponse(useCaseResp)
	return response, nil
}
