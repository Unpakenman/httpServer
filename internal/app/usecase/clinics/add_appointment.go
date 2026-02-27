package clinics

import (
	"context"
	"fmt"
	pgclient "httpServer/internal/app/client/pg"
	localerrors "httpServer/internal/app/errors"
	"httpServer/internal/app/provider"
	"httpServer/internal/app/provider/models"
	"time"
)

type AddAppointmentRequest struct {
	ClinicId   int64
	PatientId  int64
	EmployeeId int64
	StartAt    time.Time
	EndAt      time.Time
	Comment    string
}

type AddAppointmentResponse struct {
	AppointmentId int64
}

func (u *clinicsUseCase) AddAppointment(
	ctx context.Context,
	req AddAppointmentRequest,
) (*AddAppointmentResponse, localerrors.Error) {
	now := time.Now()
	if req.StartAt.Before(now) {
		err := fmt.Errorf("start at %v is before now %v", req.StartAt, now)
		return nil, localerrors.NewInternalErr(err)
	}
	var appointmentResp models.Appointments
	txErr := u.provider.WithTransaction(ctx, func(ctx context.Context, transaction pgclient.Transaction) error {
		existId, err := u.provider.CheckClinicEmployee(ctx, nil, req.ClinicId, req.EmployeeId)
		if err != nil {
			return err
		}
		if existId == nil {
			return fmt.Errorf("Employee not found in clinic")
		}
		appointmentResp, err = u.provider.AddAppointment(
			ctx,
			nil,
			provider.CreateAddAppointmentRequest{
				ClinicId:   req.ClinicId,
				PatientId:  req.PatientId,
				EmployeeId: req.EmployeeId,
				StartAt:    req.StartAt,
				EndAt:      req.EndAt,
				Comment:    req.Comment,
			})
		if err != nil {
			return localerrors.NewInternalErr(err)
		}
		return nil
	})

	if txErr != nil {
		return nil, localerrors.NewInternalErr(txErr)
	}
	return &AddAppointmentResponse{
		AppointmentId: appointmentResp.AppointmentId,
	}, nil
}
