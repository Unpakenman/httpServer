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
	ClinicId    int64
	PatientId   int64
	EmployeeId  int64
	StartAt     time.Time
	ServicesIDS []int64
	Comment     string
}

type AddAppointmentResponse struct {
	AppointmentId int64
}

func (u *clinicsUseCase) AddAppointment(
	ctx context.Context,
	req AddAppointmentRequest,
) (*AddAppointmentResponse, localerrors.Error) {
	//бизнесс логика же?
	now := time.Now()
	if req.StartAt.Before(now) {
		err := fmt.Errorf("start at %v is before now %v", req.StartAt, now)
		return nil, localerrors.NewInternalErr(err)
	}
	var appointmentResp models.Appointments
	txErr := u.provider.WithTransaction(ctx, func(ctx context.Context, transaction pgclient.Transaction) error {
		durationAndAmount, err := u.provider.GetDurationMinutesAndPrice(ctx, nil, req.ServicesIDS)
		if err != nil {
			return err
		}
		durationMinutes := durationAndAmount.DurationMinutes
		endDttm := req.StartAt.Add(time.Duration(durationMinutes) * time.Minute)
		totalPrice := durationAndAmount.TotalPrice
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
				EndAt:      endDttm,
				Comment:    req.Comment,
			})
		if err != nil {
			return localerrors.NewInternalErr(err)
		}
		appointmentID := appointmentResp.AppointmentId
		err = u.provider.CreateTransaction(
			ctx,
			nil,
			provider.CreateTransactionRequest{
				PatientId:     req.PatientId,
				ClinicId:      req.ClinicId,
				AppointmentId: appointmentID,
				Amount:        totalPrice,
				Discount:      float32(0),
				TotalAmount:   totalPrice,
				ServicesIds:   req.ServicesIDS,
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
