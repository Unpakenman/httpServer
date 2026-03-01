package provider

import (
	"context"
	"fmt"
	pgclient "httpServer/internal/app/client/pg"
	"httpServer/internal/app/provider/models"
)

type AppointmentsSlotsRequest struct {
	EmployeeId      int64
	AppointmentDTTM string
}

func (p *goExampleDBProvider) AppointmentsSlots(
	ctx context.Context,
	tx pgclient.Transaction,
	req AppointmentsSlotsRequest,
) ([]models.AppointmentsSlots, error) {
	var appointmentsSlotsData []models.AppointmentsSlots
	err := p.conn.NamedSelectContext(
		ctx,
		&appointmentsSlotsData,
		"AppointmentsSlots",
		nil,
		tx,
		req.EmployeeId,
		req.AppointmentDTTM)
	if err != nil {
		return appointmentsSlotsData, fmt.Errorf("AppointmentsSlots query error: %w", err)
	}
	return appointmentsSlotsData, nil
}
