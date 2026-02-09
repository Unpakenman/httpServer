package mapper

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/usecase/clinics"
)

func (m *mapper) ProtoToAddClinicRequest(req *pb.AddClinicRequest) clinics.AddClinicRequest {
	return clinics.AddClinicRequest{
		ClinicAdress: req.ClinicAddress,
		Phone:        req.Phone,
		Email:        req.Email,
		OpeningHours: req.OpeningHours,
	}
}

func (m *mapper) AddClinicResponseToProtoResponse(resp clinics.AddClinicResponse) *pb.AddClinicResponse {
	return &pb.AddClinicResponse{
		ClinicId: resp.ClinicId,
	}
}
