package clinic

import (
	"context"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/usecase/clinics"
)

func (s *ServerClinic) AddAppointment(
	ctx context.Context,
	req *pb.AddAppointmentRequest,
) (*pb.AddAppointmentResponse, error) {
	resp := clinics.AddAppointmentResponse{
		AppointmentId: int64(1),
	}
	response := s.mapper.AddAppointmentResponseToProtoResponse(resp)
	return response, nil
}
