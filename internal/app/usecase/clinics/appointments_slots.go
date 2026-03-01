package clinics

import (
	"context"
	localerrors "httpServer/internal/app/errors"
	"httpServer/internal/app/provider"
)

type AppointmentsSlotsRequest struct {
	EmployeeID      int64
	AppointmentDTTM string
}

type AppointmentsSlotsItems struct {
	StartAt string
	EndAt   string
}

type AppointmentsSlotsResponse struct {
	AppointmentsSlotsItems []AppointmentsSlotsItems
}

func (u *clinicsUseCase) AppointmentsSlots(
	ctx context.Context,
	req AppointmentsSlotsRequest,
) (*AppointmentsSlotsResponse, localerrors.Error) {
	appointmentsSlotsFromDB, err := u.provider.AppointmentsSlots(
		ctx,
		nil,
		provider.AppointmentsSlotsRequest{
			EmployeeId:      req.EmployeeID,
			AppointmentDTTM: req.AppointmentDTTM,
		})
	if err != nil {
		return nil, localerrors.NewInternalErr(err)
	}
	appointmentsSlotsList := make([]AppointmentsSlotsItems, 0, len(appointmentsSlotsFromDB))
	for _, appointment := range appointmentsSlotsFromDB {
		appointmentsSlotsList = append(appointmentsSlotsList, AppointmentsSlotsItems{
			StartAt: appointment.StartAt,
			EndAt:   appointment.EndAt,
		})
	}
	return &AppointmentsSlotsResponse{
		AppointmentsSlotsItems: appointmentsSlotsList,
	}, nil
}
