package provider

import (
	"context"
	"fmt"
	pgclient "httpServer/internal/app/client/pg"
	"httpServer/internal/app/provider/models"
	"time"
)

type CreatePatient struct {
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

func (p *goExampleDBProvider) CreatePatient(
	ctx context.Context,
	tx pgclient.Transaction,
	data CreatePatient,
) (models.Patients, error) {

	var patientData models.Patients

	err := p.conn.NamedGetContext(
		ctx,
		&patientData,
		"CreatePatient",
		nil,
		tx,
		data.FirstName,
		data.LastName,
		data.MiddleName,
		data.DocumentType,
		data.DocumentSeries,
		data.DocumentNumber,
		data.Sex,
		data.BirthDate,
		data.PhoneNumber,
		data.Email,
	)
	if err != nil {
		return patientData, fmt.Errorf("CreatePatient query %w", err)
	}

	return patientData, nil
}
