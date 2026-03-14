package clinic

import (
	"context"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	localerrors "httpServer/internal/app/errors"
)

func (s *ServerClinic) AddEmployee(
	ctx context.Context,
	req *pb.AddEmployeeRequest,
) (*pb.AddEmployeeResponse, error) {
	if errs := s.validator.AddEmployee(req); errs != nil {
		err := localerrors.NewInvalidArgumentErr(*errs)
		s.log.InfoCtx(ctx, "AddEmployee validation error: ", err.Error())
		return nil, s.mapper.ResultErrorToProtoError(err)
	}
	useCaseReq := s.mapper.ProtoToAddEmployeeRequest(req)
	useCaseResp, err := s.clinicUseCase.AddEmployee(ctx, useCaseReq)
	if err != nil {
		s.log.ErrorCtx(ctx, err, "AddEmployee UseCaseError")
		return nil, err

	}
	response := s.mapper.AddEmployeeResponseToProtoResponse(useCaseResp)
	return response, nil
}
