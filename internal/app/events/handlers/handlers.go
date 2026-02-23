package handlers

import (
	"httpServer/internal/app/events"
	logger "httpServer/internal/app/log"
	"httpServer/internal/app/usecase/clinics"
)

func NewHandlerList(
	logger logger.LogClient,
	clinicUseCase clinics.UseCase,
) []events.Handler {
	return []events.Handler{
		NewCreateNewServiceHandler(logger, clinicUseCase),
	}
}
