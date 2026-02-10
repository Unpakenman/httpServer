package validator

import (
	"fmt"
	pb "github.com/Unpakenman/protos/gen/go/sso/rpc"
)

func (v *validator) AddClinic(req *pb.AddClinicRequest) error {
	if req.ClinicAddress == "" || req.Email == "" || req.Phone == "" || req.OpeningHours == "" {
		return fmt.Errorf("Проверь что заполнены все поля")
	}
	return nil
}
