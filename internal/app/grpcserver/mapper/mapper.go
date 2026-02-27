package mapper

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"google.golang.org/grpc/codes"
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
	ResultErrorToProto(code codes.Code, resultError localerrors.Error) error
	AddEmployeeResponseToProtoResponse(resp *clinics.AddEmployeeResponse) *pb.AddEmployeeResponse
}

func New() Mapper { return &mapper{} }
