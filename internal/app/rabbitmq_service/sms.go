package rabbitmq_service

import (
	"context"
	"encoding/json"
	"fmt"
)

type SendSmsMessage struct {
	ClinicID     string
	PatientID    string
	PatientName  string
	PatientPhone string
	Text         string
}

type SmsMessage struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

func (s *rabbitmqService) SendSmsMessage(ctx context.Context, req SendSmsMessage) error {
	smsMessageBody, err := json.Marshal(SmsMessage{
		From: req.ClinicID,
		To:   req.PatientPhone,
		Text: req.Text,
	})
	if err != nil {
		return fmt.Errorf("marshal sms message: %w", err)
	}
	if err := s.PublishMessage(ctx, s.smsPublisher, smsMessageBody); err != nil {
		return fmt.Errorf("publish sms message: %w", err)
	}
	return nil
}
