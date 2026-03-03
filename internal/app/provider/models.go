package provider

import (
	"context"
	pgclient "httpServer/internal/app/client/pg"
	"httpServer/internal/app/provider/models"
)

//go:generate ../../../../bin/mockery --with-expecter --case=underscore --name=GoExampleProvider

type GoExampleProvider interface {
	WithTransaction(ctx context.Context, fn func(context.Context, pgclient.Transaction) error) error

	CreatePatient(
		ctx context.Context,
		tx pgclient.Transaction,
		data CreatePatientRequest,
	) (models.Patients, error)
	CreateClinic(
		ctx context.Context,
		tx pgclient.Transaction,
		data CreateClinicRequest,
	) (models.Clinic, error)
	AddEmployee(
		ctx context.Context,
		tx pgclient.Transaction,
		data CreateAddEmployeeRequest,
	) (models.Employees, error)
	AddAppointment(
		ctx context.Context,
		tx pgclient.Transaction,
		data CreateAddAppointmentRequest,
	) (models.Appointments, error)
	CheckClinicEmployee(
		ctx context.Context,
		tx pgclient.Transaction,
		clinicId int64,
		employeeId int64,
	) (*models.CheckClinicEmployee, error)
	CreateNewService(
		ctx context.Context,
		tx pgclient.Transaction,
		data CreateNewServiceRequest,
	) error
	AppointmentList(
		ctx context.Context,
		tx pgclient.Transaction,
	) ([]models.AppointmentList, error)
	AppointmentsSlots(
		ctx context.Context,
		tx pgclient.Transaction,
		req AppointmentsSlotsRequest,
	) ([]models.AppointmentsSlots, error)
	GetDurationMinutesAndPrice(
		ctx context.Context,
		tx pgclient.Transaction,
		servicesIds []int64,
	) (*models.DurationAmount, error)
	CreateTransaction(
		ctx context.Context,
		tx pgclient.Transaction,
		data CreateTransactionRequest,
	) error
}
