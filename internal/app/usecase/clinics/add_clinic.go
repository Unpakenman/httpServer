package clinics

import (
	"context"
	"httpServer/internal/app/provider"
	"strconv"
)

type AddClinicRequest struct {
	ClinicAdress string
	Phone        string
	Email        string
	OpeningHours string
}

type AddClinicResponse struct {
	ClinicId string
}

func (u *clinicsUseCase) AddClinic(
	ctx context.Context,
	req AddClinicRequest) (AddClinicResponse, error) {
	result, err := u.provider.CreateClinic(ctx, nil, provider.CreateClinicRequest{
		ClinicAddress: req.ClinicAdress,
		Phone:         req.Phone,
		Email:         req.Email,
		OpeningHours:  req.OpeningHours,
	})
	if err != nil {
		return AddClinicResponse{
			ClinicId: "0",
		}, err
	}
	clinicIdResp := strconv.FormatInt(result.ClinicID, 10)
	return AddClinicResponse{
		ClinicId: clinicIdResp,
	}, nil
}
