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
	var appointmentResp models.Appointments
	servicesInfo, err := u.provider.GetDurationMinutesAndPrice(ctx, nil, req.ServicesIDS)
	if err != nil {
		return nil, localerrors.NewInternalErr(err)
	}

	existId, err := u.provider.CheckClinicEmployee(ctx, nil, req.ClinicId, req.EmployeeId)
	if err != nil {
		return nil, localerrors.NewInternalErr(err)
	}

	if existId == nil {
		return nil, localerrors.NewBadRequestErr(fmt.Errorf("The doctor does not work at this clinic"))
	}

	txErr := u.provider.WithTransaction(ctx, func(ctx context.Context, tx pgclient.Transaction) error {
		appointmentResp, err = u.provider.AddAppointment(ctx, tx, provider.CreateAddAppointmentRequest{
			ClinicId:   req.ClinicId,
			PatientId:  req.PatientId,
			EmployeeId: req.EmployeeId,
			StartAt:    req.StartAt,
			EndAt:      req.StartAt.Add(time.Duration(servicesInfo.DurationMinutes) * time.Minute),
			Comment:    req.Comment,
		})
		if err != nil {
			return localerrors.NewInternalErr(err)
		}

		appointmentID := appointmentResp.AppointmentId
		if err := u.provider.CreateAppointmentsServices(ctx, tx, appointmentID, req.ServicesIDS); err != nil {
			return localerrors.NewInternalErr(err)
		}

		if err := u.provider.CreateTransaction(ctx, tx, provider.CreateTransactionRequest{
			PatientId:     req.PatientId,
			ClinicId:      req.ClinicId,
			AppointmentId: appointmentID,
			Amount:        servicesInfo.TotalPrice,
			Discount:      float32(0),
			TotalAmount:   servicesInfo.TotalPrice,
			ServicesIds:   req.ServicesIDS,
		}); err != nil {
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
