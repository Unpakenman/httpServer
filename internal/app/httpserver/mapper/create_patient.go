package mapper

import (
	"httpServer/internal/app/httpserver/models"
	"httpServer/internal/app/usecase/clinics"
	"strconv"
)

func (m *mapper) HttpToCreatePayinRequest(req models.CreatePatientRequest) clinics.CreatePatientRequest {
	return clinics.CreatePatientRequest{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		MiddleName:     req.MiddleName,
		DocumentType:   req.DocumentType,
		DocumentSeries: req.DocumentSeries,
		DocumentNumber: req.DocumentNumber,
		Sex:            req.Sex,
		BirthDate:      req.BirthDate,
		PhoneNumber:    req.PhoneNumber,
		Email:          req.Email,
	}
}

func (m *mapper) CreatePatientToHttp(response clinics.CreatePatientResponse) models.CreatePatientResponse {
	patientID := strconv.FormatInt(response.PatientID, 10)

	return models.CreatePatientResponse{
		PatientID: &patientID,
	}
}
