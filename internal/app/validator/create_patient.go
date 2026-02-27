package validator

import (
	"github.com/gobuffalo/validate"
	localerrors "httpServer/internal/app/errors"
	"httpServer/internal/app/httpserver/models"
	"regexp"
)

func (v *validator) CreatePatient(data models.CreatePatientRequest) *[]localerrors.FieldViolation {
	checks := []validate.Validator{
		&StringLenGreaterThenValidator{
			Name:  "first_name",
			Field: data.FirstName,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "last_name",
			Field: data.LastName,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "email",
			Field: data.Email,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "phone",
			Field: data.PhoneNumber,
			Min:   1,
		},
		&IsGreaterThanValidator[int32]{
			Name:  "document_number",
			Field: data.DocumentNumber,
			Min:   1,
		},
		&IsGreaterThanValidator[int32]{
			Name:  "special_number",
			Field: data.DocumentSeries,
			Min:   1,
		},
	}
	errors := validate.Validate(checks...)
	var fieldErr = []localerrors.FieldViolation{
		{
			Field:       "phone_number",
			Description: "phone_number does not match the pattern",
		},
	}
	var phoneRegex = regexp.MustCompile(`^(?:\+7|8)?9\d{2}\d{7}$`)
	if !phoneRegex.MatchString(data.PhoneNumber) {
		return &fieldErr
	}
	return FormatValidateErrors(errors)
}
