package mapper

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/usecase/clinics"
)

func (m *mapper) ProtoToAppointmentsSlots(req *pb.AppointmentsSlotsRequest) clinics.AppointmentsSlotsRequest {
	return clinics.AppointmentsSlotsRequest{
		EmployeeID:      req.EmployeeId,
		AppointmentDTTM: req.AppointmentDate,
	}
}

func (m *mapper) AppointmentsSlotsToProto(resp *clinics.AppointmentsSlotsResponse) *pb.AppointmentsSlotsResponse {
	slots := make([]*pb.AppointmentsSlotsResponse_Appointments, len(resp.AppointmentsSlotsItems))
	for i, item := range resp.AppointmentsSlotsItems {
		slots[i] = &pb.AppointmentsSlotsResponse_Appointments{
			SlotStart: item.StartAt,
			SlotEnd:   item.EndAt,
		}
	}
	return &pb.AppointmentsSlotsResponse{
		Appointments: slots,
	}
}
