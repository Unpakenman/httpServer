package provider

import (
	"context"
	"fmt"
	pgclient "httpServer/internal/app/client/pg"
	"httpServer/internal/app/provider/models"
	"time"
)

type CreateAddAppointmentRequest struct {
	ClinicId   int64
	PatientId  int64
	EmployeeId int64
	StartAt    time.Time
	EndAt      time.Time
	Comment    string
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
		data.StartAt,
		data.EndAt,
		data.Comment,
	)
	if err != nil {
		return appointmentData, fmt.Errorf("create appointment query error: %w", err)
	}
	return appointmentData, nil
}

func (p *goExampleDBProvider) CheckClinicEmployee(
	ctx context.Context,
	tx pgclient.Transaction,
	clinicId int64,
	employeeId int64,
) (*models.CheckClinicEmployee, error) {
	var id models.CheckClinicEmployee
	err := p.conn.NamedGetContext(
		ctx,
		&id,
		"CheckClinicEmployee",
		nil,
		tx,
		clinicId,
		employeeId)
	if err != nil {
		return &id, fmt.Errorf("check clinic employee query error: %w", err)
	}
	return &id, nil
}
