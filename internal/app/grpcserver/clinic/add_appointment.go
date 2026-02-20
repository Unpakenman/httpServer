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
	if errs := s.validator.AddAppointment(req); errs != nil {
		err := localerrors.NewInvalidArgumentErr(*errs)
		s.log.InfoCtx(ctx, "AddClinic validation error: ", err)
		return nil, s.mapper.ResultErrorToProto(err)
	}
	useCaseReq, err := s.mapper.ProtoToAddAppointmentRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	useCaseResp, err := s.clinicUseCase.AddAppointment(ctx, useCaseReq)
	if err != nil {
		s.log.ErrorCtx(ctx, err)
		return nil, err
	}
	response := s.mapper.AddAppointmentResponseToProtoResponse(useCaseResp)
	return response, nil
}
