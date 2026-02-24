package provider

import (
	"context"
	pgclient "httpServer/internal/app/client/pg"
	"httpServer/internal/app/provider/models"
)

func (p *goExampleDBProvider) AppointmentList(
	ctx context.Context,
	tx pgclient.Transaction,
) ([]models.AppointmentList, error) {
	var appointmentsData []models.AppointmentList
	err := p.conn.NamedSelectContext(
		ctx,
		&appointmentsData,
		"AppointmentsList",
		nil,
		tx)
	if err != nil {
		return nil, err
	}
	return appointmentsData, nil
}
