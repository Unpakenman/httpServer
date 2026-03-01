package validator

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"github.com/gobuffalo/validate"
	localerrors "httpServer/internal/app/errors"
)

func (v *validator) AppointmentsSlots(req *pb.AppointmentsSlotsRequest) *[]localerrors.FieldViolation {
	checks := []validate.Validator{
		&IsGreaterThanValidator[int64]{
			Name:  "employee_id",
			Field: req.EmployeeId,
			Min:   1,
		},

		&StringLenGreaterThenValidator{
			Name:  "appointment_date",
			Field: req.AppointmentDate,
			Min:   1,
		},
	}
	errors := validate.Validate(checks...)
	return FormatValidateErrors(errors)
}
