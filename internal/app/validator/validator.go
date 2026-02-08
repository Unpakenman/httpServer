package validator

import "httpServer/internal/app/httpserver/models"

type validator struct{}

type Validator interface {
	CreatePatient(data models.CreatePatientRequest) error
}

func New() Validator {
	return &validator{}
}
