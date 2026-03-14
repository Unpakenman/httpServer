package mapper

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/usecase/clinics"
	"time"
)

func (m *mapper) ProtoToAddAppointmentRequest(req *pb.AddAppointmentRequest) (clinics.AddAppointmentRequest, error) {
	startAt, err := time.Parse("2006-01-02T15:04:05Z", req.StartAt)
	if err != nil {
		return clinics.AddAppointmentRequest{}, err
	}
	return clinics.AddAppointmentRequest{
		ClinicId:    req.ClinicId,
		PatientId:   req.PatientId,
		EmployeeId:  req.EmployeeId,
		StartAt:     startAt,
		ServicesIDS: req.ServicesIds,
		Comment:     req.Comment,
	}, nil
}

func (m *mapper) AddAppointmentResponseToProtoResponse(resp *clinics.AddAppointmentResponse) *pb.AddAppointmentResponse {
	return &pb.AddAppointmentResponse{
		AppointmentId: resp.AppointmentId,
	}
}
