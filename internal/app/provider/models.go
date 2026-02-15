package provider

import (
	"context"
	pgclient "httpServer/internal/app/client/pg"
	"httpServer/internal/app/provider/models"
)

//go:generate ../../../../bin/mockery --with-expecter --case=underscore --name=GoExampleProvider

type GoExampleProvider interface {
	BeginTransaction() (pgclient.Transaction, error)
	CommitTransaction(tx pgclient.Transaction) error
	RollbackTransaction(tx pgclient.Transaction)

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
}
