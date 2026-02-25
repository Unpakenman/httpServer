package jobs

import (
	"context"
	"httpServer/internal/app/usecase/clinics"
)

type Job interface {
	GetName() string
	Run(ctx context.Context) error
}

func NewJobList(
	clinicssUsecase clinics.UseCase,
) []Job {
	return []Job{
		NewAppointmentsList(clinicssUsecase),
	}
}
