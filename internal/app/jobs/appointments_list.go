package jobs

import (
	"context"
	"httpServer/internal/app/usecase/clinics"
)

const (
	appointmentListJobName = "appointment_list"
)

type ApointmentsListJob struct {
	usecase clinics.UseCase
}

func NewAppointmentsList(usecase clinics.UseCase) *ApointmentsListJob {
	return &ApointmentsListJob{
		usecase: usecase,
	}
}

func (j *ApointmentsListJob) Run(ctx context.Context) error {
	err := j.usecase.AppointmentList(ctx)
	return err
}

func (j *ApointmentsListJob) GetName() string { return appointmentListJobName }
