package mapper

import (
	"httpServer/internal/app/httpserver/models"
	"httpServer/internal/app/usecase/clinics"
)

type mapper struct{}
type Mapper interface {
	HttpToCreatePayinRequest(req models.CreatePatientRequest) clinics.CreatePatientRequest
}

func New() Mapper { return &mapper{} }
