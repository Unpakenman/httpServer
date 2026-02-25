package rabbitmq_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type PublishMessage struct {
	EventID    int64                     `json:"event_id"`
	Name       string                    `json:"name"`
	Uid        string                    `json:"uid"`
	CreateDttm string                    `json:"create_dttm"`
	Data       []AppointmentsListMessage `json:"data"`
}

type AppointmentsListMessage struct {
	PatientFirstName       string          `json:"patient_first_name"`
	PatientLastName        string          `json:"patient_last_name"`
	DoctorFirstName        string          `json:"doctor_first_name"`
	DoctorLastName         string          `json:"doctor_last_name"`
	ClinicAddress          string          `json:"clinic_address"`
	Price                  decimal.Decimal `json:"price"`
	AppointmentDTTM        string          `json:"appointment_dttm"`
	AppointmentDescription string          `json:"appointment_description"`
}

func (s *rabbitmqService) SendAppointmentsMessage(ctx context.Context, req []AppointmentsListMessage) error {
	smsMessageBody, err := json.Marshal(PublishMessage{
		EventID:    0,
		Name:       "appointments_list",
		Uid:        uuid.NewString(),
		CreateDttm: time.Now().Format(time.RFC3339),
		Data:       req,
	})
	if err != nil {
		return fmt.Errorf("marshal sms message: %w", err)
	}
	if err := s.PublishMessage(ctx, s.smsPublisher, smsMessageBody); err != nil {
		return fmt.Errorf("publish sms message: %w", err)
	}
	return nil
}
