package provider

import (
	"context"
	"fmt"
	pgclient "httpServer/internal/app/client/pg"
	"httpServer/internal/app/provider/models"
)

type CreateClinicRequest struct {
	ClinicAddress string
	Phone         string
	Email         string
	OpeningHours  string
}

func (p *goExampleDBProvider) CreateClinic(
	ctx context.Context,
	tx pgclient.Transaction,
	data CreateClinicRequest,
) (models.Clinic, error) {
	var clinicData models.Clinic
	err := p.conn.NamedGetContext(
		ctx,
		&clinicData,
		"CreateClinic",
		nil,
		tx,
		data.ClinicAddress,
		data.Phone,
		data.Email,
		data.OpeningHours)
	if err != nil {
		return clinicData, fmt.Errorf("CreatePatient query %w", err)
	}
	return clinicData, nil
}
