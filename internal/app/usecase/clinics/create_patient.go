package clinics

import (
	"context"
	"httpServer/internal/app/provider"
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
	PatientID int64
}

func (u *clinicsUseCase) CreatePatient(
	ctx context.Context,
	req CreatePatientRequest,
) (CreatePatientResponse, error) {
	result, err := u.provider.CreatePatient(ctx, nil, provider.CreatePatientRequest{
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
	})
	if err != nil {
		return CreatePatientResponse{}, err
	}

	return CreatePatientResponse{
		PatientID: result.PatientID,
	}, nil
}
