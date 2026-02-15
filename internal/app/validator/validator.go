package validator

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	localerrors "httpServer/internal/app/errors"
	"httpServer/internal/app/httpserver/models"
)

type validator struct{}

type Validator interface {
	CreatePatient(data models.CreatePatientRequest) error
	AddClinic(req *pb.AddClinicRequest) *[]localerrors.FieldViolation
	AddEmployee(req *pb.AddEmployeeRequest) error
}

func New() Validator {
	return &validator{}
}
