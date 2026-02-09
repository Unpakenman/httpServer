package clinic

import (
	"context"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/usecase/clinics"
)

func (s *ServerClinic) AddEmployee(
	ctx context.Context,
	req *pb.AddEmployeeRequest,
) (*pb.AddEmployeeResponse, error) {
	resp := clinics.AddEmployeeResponse{EmployeeId: int64(1)}
	response := s.mapper.AddEmployeeResponseToProtoResponse(resp)
	return response, nil
}
