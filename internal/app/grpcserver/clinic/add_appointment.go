package clinic

import (
	"context"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	localerrors "httpServer/internal/app/errors"
)

func (s *ServerClinic) AddAppointment(
	ctx context.Context,
	req *pb.AddAppointmentRequest,
) (*pb.AddAppointmentResponse, error) {
	if errs := s.validator.AddAppointment(req); errs != nil {
		err := localerrors.NewInvalidArgumentErr(*errs)
		s.log.InfoCtx(ctx, "AddClinic validation error: ", err.Error())
		return nil, s.mapper.ResultErrorToProtoError(err)
	}
	useCaseReq, err := s.mapper.ProtoToAddAppointmentRequest(req)
	if err != nil {
		return nil, s.mapper.ResultErrorToProtoError(localerrors.NewInternalErr(err))
	}
	useCaseResp, err := s.clinicUseCase.AddAppointment(ctx, useCaseReq)
	if err != nil {
		s.log.ErrorCtx(ctx, err)
		return nil, s.mapper.ResultErrorToProtoError(localerrors.NewInternalErr(err))
	}
	response := s.mapper.AddAppointmentResponseToProtoResponse(useCaseResp)
	return response, nil
}
