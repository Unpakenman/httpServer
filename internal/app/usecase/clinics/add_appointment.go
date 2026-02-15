package clinics

import (
	"context"
	localerrors "httpServer/internal/app/errors"
	"httpServer/internal/app/provider"
	"time"
)

type AddAppointmentRequest struct {
	ClinicId        int64
	PatientId       int64
	EmployeeId      int64
	AppointmentDTTM string
	Comment         string
}

type AddAppointmentResponse struct {
	AppointmentId int64
}

func (u *clinicsUseCase) AddAppointment(
	ctx context.Context,
	req AddAppointmentRequest) (*AddAppointmentResponse, localerrors.Error) {
	appointmentDttm, err := time.Parse("2006-01-02T15:04:05Z", req.AppointmentDTTM)
	if err != nil {
		u.logger.ErrorContext(ctx, "Failed to parse appointmentDttm")
		return &AddAppointmentResponse{}, localerrors.NewInternalErr(err)
	}
	result, err := u.provider.AddAppointment(
		ctx,
		nil,
		provider.CreateAddAppointmentRequest{
			ClinicId:        req.ClinicId,
			PatientId:       req.PatientId,
			EmployeeId:      req.EmployeeId,
			AppointmentDttm: appointmentDttm,
			Comment:         req.Comment,
		})
	if err != nil {
		u.logger.ErrorContext(ctx, "Failed to add appointment")
		return nil, localerrors.NewInternalErr(err)
	}
	return &AddAppointmentResponse{
		AppointmentId: result.AppointmentId,
	}, nil
}
