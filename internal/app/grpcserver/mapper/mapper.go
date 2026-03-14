package mapper

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	localerrors "httpServer/internal/app/errors"
	"httpServer/internal/app/usecase/clinics"
)

type mapper struct{}

type Mapper interface {
	ProtoToAddAppointmentRequest(req *pb.AddAppointmentRequest) (clinics.AddAppointmentRequest, error)
	AddAppointmentResponseToProtoResponse(resp *clinics.AddAppointmentResponse) *pb.AddAppointmentResponse
	ProtoToAddClinicRequest(req *pb.AddClinicRequest) clinics.AddClinicRequest
	AddClinicResponseToProtoResponse(resp *clinics.AddClinicResponse) *pb.AddClinicResponse
	ProtoToAddEmployeeRequest(req *pb.AddEmployeeRequest) clinics.AddEmployeeRequest
	ResultErrorToProtoError(resultError localerrors.Error) error
	AddEmployeeResponseToProtoResponse(resp *clinics.AddEmployeeResponse) *pb.AddEmployeeResponse
	ProtoToAppointmentsSlots(req *pb.AppointmentsSlotsRequest) clinics.AppointmentsSlotsRequest
	AppointmentsSlotsToProto(resp *clinics.AppointmentsSlotsResponse) *pb.AppointmentsSlotsResponse
}

func New() Mapper { return &mapper{} }
