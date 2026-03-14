package clinic

import (
	"context"
	"fmt"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	localerrors "httpServer/internal/app/errors"
)

func (s *ServerClinic) AppointmentsSlotsRequest(
	ctx context.Context,
	req *pb.AppointmentsSlotsRequest,
) (*pb.AppointmentsSlotsResponse, error) {
	if errs := s.validator.AppointmentsSlots(req); errs != nil {
		fmt.Println(errs)
		err := localerrors.NewInvalidArgumentErr(*errs)
		s.log.InfoCtx(ctx, "AppointmentsSlots validation error: ", err.Error())
		return nil, s.mapper.ResultErrorToProtoError(err)
	}
	useCaseReq := s.mapper.ProtoToAppointmentsSlots(req)
	useCaseResp, err := s.clinicUseCase.AppointmentsSlots(ctx, useCaseReq)
	if err != nil {
		s.log.ErrorCtx(ctx, err)
		return nil, s.mapper.ResultErrorToProtoError(localerrors.NewInternalErr(err))
	}
	response := s.mapper.AppointmentsSlotsToProto(useCaseResp)
	return response, nil
}
