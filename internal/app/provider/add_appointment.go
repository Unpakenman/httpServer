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

func (p *goExampleDBProvider) GetDurationMinutesAndPrice(
	ctx context.Context,
	tx pgclient.Transaction,
	servicesIds []int64,
) (*models.DurationAmount, error) {
	var durationAmount models.DurationAmount
	err := p.conn.NamedGetContext(
		ctx,
		&durationAmount,
		"GetDurationMinutesByServicesIds",
		nil,
		tx,
		servicesIds,
	)
	if err != nil {
		return nil, fmt.Errorf("get duration minutes query error: %w", err)
	}
	return &durationAmount, nil
}

type CreateTransactionRequest struct {
	PatientId     int64
	ClinicId      int64
	AppointmentId int64
	Amount        float32
	Discount      float32
	TotalAmount   float32
	ServicesIds   []int64
}

func (p *goExampleDBProvider) CreateTransaction(
	ctx context.Context,
	tx pgclient.Transaction,
	data CreateTransactionRequest,
) error {
	_, err := p.conn.Exec(
		ctx,
		"CreateTransaction",
		nil,
		tx,
		data.PatientId,
		data.ClinicId,
		data.AppointmentId,
		data.Amount,
		data.Discount,
		data.TotalAmount,
		data.ServicesIds)
	if err != nil {
		return fmt.Errorf("create transaction query error: %w", err)
	}
	return nil
}
