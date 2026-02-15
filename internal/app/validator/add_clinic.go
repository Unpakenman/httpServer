package validator

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"github.com/gobuffalo/validate"
	localerrors "httpServer/internal/app/errors"
)

func (v *validator) AddClinic(req *pb.AddClinicRequest) *[]localerrors.FieldViolation {
	checks := []validate.Validator{
		&StringLenGreaterThenValidator{
			Name:  "clinic_address",
			Field: req.ClinicAddress,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "email,",
			Field: req.Email,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "phone",
			Field: req.Phone,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "opening_hours",
			Field: req.OpeningHours,
			Min:   1,
		},
	}
	errors := validate.Validate(checks...)
	return FormatValidateErrors(errors)
}
