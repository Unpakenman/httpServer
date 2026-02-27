package clinics

import (
	"context"
	localerrors "httpServer/internal/app/errors"
	"httpServer/internal/app/rabbitmq_service"
)

func (u *clinicsUseCase) AppointmentList(ctx context.Context) localerrors.Error {
	appointmentsList, err := u.provider.AppointmentList(ctx, nil)
	if err != nil {
		return localerrors.NewInternalErr(err)
	}
	appointmentListPublish := make([]rabbitmq_service.AppointmentsListMessage, 0, len(appointmentsList))
	for _, appointment := range appointmentsList {
		aplist := rabbitmq_service.AppointmentsListMessage{
			PatientFirstName:       appointment.PatientName,
			PatientLastName:        appointment.PatientLastName,
			DoctorFirstName:        appointment.DoctorName,
			DoctorLastName:         appointment.DoctorLastName,
			ClinicAddress:          appointment.ClinicAdress,
			Price:                  appointment.Price,
			StartAt:                appointment.StartAt,
			AppointmentDescription: appointment.AppointmentDescription,
		}
		appointmentListPublish = append(appointmentListPublish, aplist)
	}
	errSend := u.rmqService.SendAppointmentsMessage(ctx, appointmentListPublish)
	if errSend != nil {
		return localerrors.NewInternalErr(errSend)
	}
	return nil
}
