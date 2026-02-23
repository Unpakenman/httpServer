package clinics

import (
	"context"
	"github.com/shopspring/decimal"
	localerrors "httpServer/internal/app/errors"
	"httpServer/internal/app/provider"
)

type CreateNewServiceRequest struct {
	Name            string
	Description     string
	Price           decimal.Decimal
	DurationMinutes int64
	IsActive        bool
}

func (u *clinicsUseCase) CreateNewService(
	ctx context.Context,
	request CreateNewServiceRequest,
) localerrors.Error {
	err := u.provider.CreateNewService(ctx, nil, provider.CreateNewServiceRequest{
		Name:            request.Name,
		Description:     request.Description,
		Price:           request.Price,
		DurationMinutes: request.DurationMinutes,
		IsActive:        request.IsActive,
	})
	if err != nil {
		return localerrors.NewInternalErr(err)
	}
	return nil
}
