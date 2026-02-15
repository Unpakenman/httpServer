package mapper

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/usecase/clinics"
)

func (m *mapper) ProtoToAddAppointmentRequest(req *pb.AddAppointmentRequest) clinics.AddAppointmentRequest {
	return clinics.AddAppointmentRequest{
		ClinicId:        req.ClinicId,
		PatientId:       req.PatientId,
		EmployeeId:      req.EmployeeId,
		AppointmentDTTM: req.AppointmentDttm,
		Comment:         req.Comment,
	}
}

func (m *mapper) AddAppointmentResponseToProtoResponse(resp *clinics.AddAppointmentResponse) *pb.AddAppointmentResponse {
	return &pb.AddAppointmentResponse{
		AppointmentId: resp.AppointmentId,
	}
}
