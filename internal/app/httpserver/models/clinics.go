package models

import "time"

type CreatePatientRequest struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	MiddleName     *string   `json:"middle_name"`
	DocumentType   int32     `json:"document_type"`
	DocumentSeries int32     `json:"document_series"`
	DocumentNumber int32     `json:"document_number"`
	Sex            string    `json:"sex"`
	BirthDate      time.Time `json:"birth_date"`
	PhoneNumber    string    `json:"phone_number"`
	Email          string    `json:"email"`
}

type CreatePatientResponse struct {
	Status    string  `json:"status"`
	PatientId *string `json:"patient_id"`
}
