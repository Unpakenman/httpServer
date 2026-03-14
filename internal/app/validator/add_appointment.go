package validator

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"github.com/gobuffalo/validate"
	localerrors "httpServer/internal/app/errors"
	"time"
)

func (v *validator) AddAppointment(req *pb.AddAppointmentRequest) *[]localerrors.FieldViolation {
	timeErr := validate.NewErrors()
	now := time.Now().UTC()
	timeRequest, err := time.Parse(time.RFC3339, req.StartAt)
	if err != nil {
		timeErr.Add("Start_at error: ", err.Error())
		return FormatValidateErrors(timeErr)
	}

	timeRequest = timeRequest.UTC()
	if timeRequest.Before(now) {
		timeErr.Add("Invalid start_at: ", "start_at is before current time")
		return FormatValidateErrors(timeErr)
	}

	checks := []validate.Validator{
		&IsGreaterThanValidator[int64]{
			Name:  "clinic_id",
			Field: req.ClinicId,
			Min:   1,
		},
		&IsGreaterThanValidator[int64]{
			Name:  "patient_id",
			Field: req.PatientId,
			Min:   1,
		},
		&IsGreaterThanValidator[int64]{
			Name:  "employee_id",
			Field: req.EmployeeId,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "start_at",
			Field: req.StartAt,
			Min:   1,
		},
		&StringIsDateIsIso8601Validator{
			Name:  "start_at",
			Field: req.StartAt,
		},
		&StringLenGreaterThenValidator{
			Name:  "comment",
			Field: req.Comment,
			Min:   1,
		},
	}
	errors := validate.Validate(checks...)
	return FormatValidateErrors(errors)
}
