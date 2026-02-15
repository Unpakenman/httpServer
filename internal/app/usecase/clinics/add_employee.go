package clinics

import (
	"context"
	"fmt"
	"httpServer/internal/app/provider"
)

type AddEmployeeRequest struct {
	RoleId           int64
	SpecializationId int64
	FirstName        string
	LastName         string
	MiddleName       *string
	BirthDate        string
	Phone            string
	Email            string
}

type AddEmployeeResponse struct {
	EmployeeId int64
}

func (u *clinicsUseCase) AddEmployee(ctx context.Context, req AddEmployeeRequest) (AddEmployeeResponse, error) {
	result, err := u.provider.AddEmployee(ctx, nil, provider.CreateAddEmployeeRequest{
		RoleId:           req.RoleId,
		SpecializationId: req.SpecializationId,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		MiddleName:       req.MiddleName,
		BirthDate:        req.BirthDate,
		Phone:            req.Phone,
		Email:            req.Email,
	})
	if err != nil {
		return AddEmployeeResponse{}, fmt.Errorf("Failed to create db request  %w", err)
	}
	return AddEmployeeResponse{
		EmployeeId: result.EmployeeID,
	}, nil
}
