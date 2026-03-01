package clinic

import (
	"context"
	"fmt"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"google.golang.org/grpc/codes"
	localerrors "httpServer/internal/app/errors"
)

func (s *ServerClinic) AppointmentsSlotsRequest(
	ctx context.Context,
	req *pb.AppointmentsSlotsRequest,
) (*pb.AppointmentsSlotsResponse, error) {
	if errs := s.validator.AppointmentsSlots(req); errs != nil {
		fmt.Println(errs)
		err := localerrors.NewInvalidArgumentErr(*errs)
		s.log.InfoCtx(ctx, "AppointmentsSlots validation error: ", err)
		return nil, s.mapper.ResultErrorToProto(codes.InvalidArgument, err)
	}
	useCaseReq := s.mapper.ProtoToAppointmentsSlots(req)
	useCaseResp, err := s.clinicUseCase.AppointmentsSlots(ctx, useCaseReq)
	if err != nil {
		s.log.ErrorCtx(ctx, err)
		return nil, s.mapper.ResultErrorToProto(codes.Internal, localerrors.NewInternalErr(err))
	}
	response := s.mapper.AppointmentsSlotsToProto(useCaseResp)
	return response, nil
}
