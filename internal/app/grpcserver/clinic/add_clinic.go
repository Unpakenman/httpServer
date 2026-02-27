package clinic

import (
	"context"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"google.golang.org/grpc/codes"
	localerrors "httpServer/internal/app/errors"
)

func (s *ServerClinic) AddClinic(
	ctx context.Context,
	req *pb.AddClinicRequest,
) (*pb.AddClinicResponse, error) {
	if errs := s.validator.AddClinic(req); errs != nil {
		err := localerrors.NewInvalidArgumentErr(*errs)
		s.log.InfoCtx(ctx, "AddClinic validation error: ", err)
		return nil, s.mapper.ResultErrorToProto(codes.InvalidArgument, err)
	}
	useCaseReq := s.mapper.ProtoToAddClinicRequest(req)
	resp, err := s.clinicUseCase.AddClinic(ctx, useCaseReq)
	if err != nil {
		s.log.ErrorCtx(ctx, err, "AddClinic UseCaseError")
		return nil, err
	}
	response := s.mapper.AddClinicResponseToProtoResponse(resp)
	return response, nil
}
