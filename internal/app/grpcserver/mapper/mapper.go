package mapper

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/usecase/clinics"
)

type mapper struct{}

type Mapper interface {
	ProtoToAddAppointmentRequest(req *pb.AddAppointmentRequest) clinics.AddAppointmentRequest
	AddAppointmentResponseToProtoResponse(resp clinics.AddAppointmentResponse) *pb.AddAppointmentResponse
	ProtoToAddClinicRequest(req *pb.AddClinicRequest) clinics.AddClinicRequest
	AddClinicResponseToProtoResponse(resp clinics.AddClinicResponse) *pb.AddClinicResponse
	ProtoToAddEmployeeRequest(req *pb.AddEmployeeRequest) clinics.AddEmployeeRequest
	AddEmployeeResponseToProtoResponse(resp clinics.AddEmployeeResponse) *pb.AddEmployeeResponse
}

func New() Mapper { return &mapper{} }
