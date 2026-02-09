package clinic

import (
	"context"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/usecase/clinics"
)

func (s *ServerClinic) AddClinic(
	ctx context.Context,
	req *pb.AddClinicRequest,
) (*pb.AddClinicResponse, error) {
	resp := clinics.AddClinicResponse{
		ClinicId: int64(1),
	}
	response := s.mapper.AddClinicResponseToProtoResponse(resp)
	return response, nil
}
