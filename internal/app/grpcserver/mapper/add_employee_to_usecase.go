package mapper

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/usecase/clinics"
)

func (m *mapper) ProtoToAddEmployeeRequest(req *pb.AddEmployeeRequest) clinics.AddEmployeeRequest {
	return clinics.AddEmployeeRequest{
		RoleId:           req.RoleId,
		SpecializationId: req.SpecializationId,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		MiddleName:       req.MiddleName,
		BirthDate:        req.Birthdate,
		Email:            req.Email,
		Phone:            req.Phone,
	}
}

func (m *mapper) AddEmployeeResponseToProtoResponse(resp clinics.AddEmployeeResponse) *pb.AddEmployeeResponse {
	return &pb.AddEmployeeResponse{
		EmployeeId: resp.EmployeeId,
	}
}
