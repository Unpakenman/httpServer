package validator

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"httpServer/internal/app/httpserver/models"
)

type validator struct{}

type Validator interface {
	CreatePatient(data models.CreatePatientRequest) error
	AddClinic(req *pb.AddClinicRequest) error
}

func New() Validator {
	return &validator{}
}
