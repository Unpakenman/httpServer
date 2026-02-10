package mapper

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/usecase/clinics"
	"log"
	"strconv"
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
	clinicIdResp, err := strconv.ParseInt(resp.ClinicId, 10, 64)
	if err != nil {
		log.Fatalf("could not convert clinic id to int: %v", err)
	}
	return &pb.AddClinicResponse{
		ClinicId: clinicIdResp,
	}
}
