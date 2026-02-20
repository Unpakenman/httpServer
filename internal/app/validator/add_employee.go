package validator

import (
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"github.com/gobuffalo/validate"
	localerrors "httpServer/internal/app/errors"
)

func (v *validator) AddEmployee(req *pb.AddEmployeeRequest) *[]localerrors.FieldViolation {
	checks := []validate.Validator{
		&IsGreaterThanValidator[int64]{
			Name:  "role_id",
			Field: req.RoleId,
			Min:   1,
		},
		&IsGreaterThanValidator[int64]{
			Name:  "specialixationa_id",
			Field: req.SpecializationId,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "first_name",
			Field: req.FirstName,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "last_name",
			Field: req.LastName,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "birth_date",
			Field: req.Birthdate,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "phone",
			Field: req.Phone,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "email",
			Field: req.Email,
			Min:   1,
		},
	}
	errors := validate.Validate(checks...)
	return FormatValidateErrors(errors)
}
