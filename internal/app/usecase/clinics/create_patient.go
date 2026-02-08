package clinics

import (
	"context"
	"httpServer/internal/app/provider"
	"strconv"
	"time"
)

type CreatePatientRequest struct {
	FirstName      string
	LastName       string
	MiddleName     *string
	DocumentType   int32
	DocumentSeries int32
	DocumentNumber int32
	Sex            string
	BirthDate      time.Time
	PhoneNumber    string
	Email          string
}

type CreatePatientResponse struct {
	PatientId string
}

func (u *clinicsUseCase) CreatePatient(
	ctx context.Context,
	req CreatePatientRequest) (CreatePatientResponse, error) {
	result, err := u.provider.CreatePatient(
		ctx,
		nil,
		provider.CreatePatient{
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
		},
	)
	if err != nil {
		return CreatePatientResponse{
			PatientId: "0",
		}, err
	}
	patientId := strconv.FormatInt(result.PatientID, 10)
	return CreatePatientResponse{
		PatientId: patientId,
	}, nil
}
