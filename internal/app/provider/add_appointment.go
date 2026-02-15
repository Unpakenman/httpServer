package provider

import (
	"context"
	"fmt"
	pgclient "httpServer/internal/app/client/pg"
	"httpServer/internal/app/provider/models"
	"time"
)

type CreateAddAppointmentRequest struct {
	ClinicId        int64
	PatientId       int64
	EmployeeId      int64
	AppointmentDttm time.Time
	Comment         string
}

func (p *goExampleDBProvider) AddAppointment(
	ctx context.Context,
	tx pgclient.Transaction,
	data CreateAddAppointmentRequest,
) (models.Appointments, error) {
	var appointmentData models.Appointments
	err := p.conn.NamedGetContext(
		ctx,
		&appointmentData,
		"CreateAppointment",
		nil,
		tx,
		data.ClinicId,
		data.PatientId,
		data.EmployeeId,
		data.AppointmentDttm,
		data.Comment)
	if err != nil {
		return appointmentData, fmt.Errorf("create appointment query error: %w", err)
	}
	return appointmentData, nil
}
