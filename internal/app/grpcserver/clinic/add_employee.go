package clinic

import (
	"context"
	"fmt"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
)

func (s *ServerClinic) AddEmployee(
	ctx context.Context,
	req *pb.AddEmployeeRequest,
) (*pb.AddEmployeeResponse, error) {
	fmt.Println("AddEmployee called")
	if err := s.validator.AddEmployee(req); err != nil {
		return nil, fmt.Errorf("Ошибка валидации  %w", err)
	}
	useCaseReq := s.mapper.ProtoToAddEmployeeRequest(req)
	useCaseResp, useCaseErr := s.clinicUseCase.AddEmployee(ctx, useCaseReq)
	if useCaseErr != nil {
		return nil, fmt.Errorf("UseCase AddEmployee error  %w", useCaseErr)
	}
	response := s.mapper.AddEmployeeResponseToProtoResponse(useCaseResp)
	return response, nil
}
