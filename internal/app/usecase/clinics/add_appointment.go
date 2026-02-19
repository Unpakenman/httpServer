package clinics

import (
	"context"
	"fmt"
	localerrors "httpServer/internal/app/errors"
	"httpServer/internal/app/provider"
	"time"
)

type AddAppointmentRequest struct {
	ClinicId        int64
	PatientId       int64
	EmployeeId      int64
	AppointmentDTTM time.Time
	Comment         string
}

type AddAppointmentResponse struct {
	AppointmentId int64
}

func (u *clinicsUseCase) AddAppointment(
	ctx context.Context,
	req AddAppointmentRequest,
) (*AddAppointmentResponse, localerrors.Error) {
	result, err := u.provider.AddAppointment(
		ctx,
		nil,
		provider.CreateAddAppointmentRequest{
			ClinicId:        req.ClinicId,
			PatientId:       req.PatientId,
			EmployeeId:      req.EmployeeId,
			AppointmentDttm: req.AppointmentDTTM,
			Comment:         req.Comment,
		})
	if err != nil {
		return nil, localerrors.NewInternalErr(fmt.Errorf("AddAppointment error: %w", err))
	}
	return &AddAppointmentResponse{
		AppointmentId: result.AppointmentId,
	}, nil
}
