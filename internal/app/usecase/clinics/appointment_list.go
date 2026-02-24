package clinics

import (
	"context"
	localerrors "httpServer/internal/app/errors"
)

type Appointments struct {
	PatientFirstName       string
	PatientLastName        string
	DoctorFirstName        string
	DoctorLastName         string
	ClinicAddress          string
	Price                  string
	AppointmentDTTM        string
	AppointmentDescription string
}

func (u *clinicsUseCase) AppointmentList(ctx context.Context) localerrors.Error {
	appointmentsList, err := u.provider.AppointmentList(ctx, nil)
	if err != nil {
		return localerrors.NewInternalErr(err)
	}
	u.logger.DebugCtx(ctx, "AppointmentList", appointmentsList)
	return nil
}
