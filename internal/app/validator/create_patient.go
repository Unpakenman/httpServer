package validator

import (
	"errors"
	"httpServer/internal/app/httpserver/models"
	"regexp"
)

func (v *validator) CreatePatient(data models.CreatePatientRequest) error {
	if data.FirstName == "" {
		return errors.New("first name is required")
	}
	if data.LastName == "" {
		return errors.New("last name is required")
	}
	if data.PhoneNumber == "" {
		return errors.New("phone number is required")
	}
	if data.Email == "" {
		return errors.New("email is required")
	}
	if data.DocumentNumber <= 0 {
		return errors.New("mistake in document number")
	}
	if data.DocumentSeries <= 0 {
		return errors.New("mistake in document series")
	}
	if data.PhoneNumber == "" || len(data.PhoneNumber) < 10 || len(data.PhoneNumber) > 12 {
		return errors.New("phone number or phone number is invalid")
	}
	if _, err := regexp.MatchString(`^(?:\+7|8)?9\d{2}\d{7}$`, data.PhoneNumber); err != nil {
		return err
	}
	/*if data.BirthDate.Before(time.Now()) {
		return errors.New("birth date is required")
	}*/
	return nil
}
