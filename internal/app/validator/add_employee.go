package validator

import (
	"fmt"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
)

func (v *validator) AddEmployee(req *pb.AddEmployeeRequest) error {
	if req.Phone == "" || req.Email == "" || req.RoleId <= 0 || req.SpecializationId <= 0 || req.LastName == "" || req.FirstName == "" || req.Birthdate == "" {
		return fmt.Errorf("Неверно заполнены поля")
	}
	return nil
}
