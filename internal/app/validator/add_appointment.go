package validator

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"github.com/gobuffalo/validate"
	localerrors "httpServer/internal/app/errors"
)

func (v *validator) AddAppointment(req *pb.AddAppointmentRequest) *[]localerrors.FieldViolation {
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
		&StringLenGreaterThenValidator{
			Name:  "comment",
			Field: req.Comment,
			Min:   1,
		},
	}
	errors := validate.Validate(checks...)
	return FormatValidateErrors(errors)
}
