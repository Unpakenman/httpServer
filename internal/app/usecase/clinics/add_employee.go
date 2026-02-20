package clinics

import (
	"context"
	localerrors "httpServer/internal/app/errors"
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

func (u *clinicsUseCase) AddEmployee(
	ctx context.Context, req AddEmployeeRequest,
) (*AddEmployeeResponse, localerrors.Error) {
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
		return nil, localerrors.NewInternalErr(err)
	}
	return &AddEmployeeResponse{
		EmployeeId: result.EmployeeID,
	}, nil
}
