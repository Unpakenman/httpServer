package handlers

import (
	"context"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"httpServer/internal/app/events"
	"httpServer/internal/app/events/handlers/models"
	logger "httpServer/internal/app/log"
	"httpServer/internal/app/usecase/clinics"
)

type CreateNewServiceHandler struct {
	logger  logger.LogClient
	usecase clinics.UseCase
}

func NewCreateNewServiceHandler(
	logger logger.LogClient,
	usecase clinics.UseCase,
) events.Handler {
	return &CreateNewServiceHandler{
		logger:  logger,
		usecase: usecase,
	}
}

func (h *CreateNewServiceHandler) Description() events.HandlerDescription {
	return events.HandlerDescription{
		Alias:       "create_new_service",
		Description: "handler for event: create_new_service",
	}
}

func (h *CreateNewServiceHandler) Handle(
	ctx context.Context,
	rawData []byte,
	meta events.Meta,
) events.HandleResult {
	var data events.EventData[models.CreateNewServiceRequest]
	err := jsoniter.Unmarshal(rawData, &data)
	if err != nil {
		err = errors.Join(err, errors.New("amqp failed unmarshal event data"))
		h.logger.ErrorCtx(ctx, err)
		return events.MakeInternalErrorHandleResult(err)
	}

	err = h.usecase.CreateNewService(ctx, clinics.CreateNewServiceRequest{
		Name:            data.Data.Name,
		Description:     data.Data.Description,
		Price:           data.Data.Price,
		DurationMinutes: data.Data.DurationMinutes,
		IsActive:        data.Data.IsActive,
	})
	if err != nil {
		err := errors.Join(err, errors.New("amqp internal error"))
		return events.MakeInternalErrorHandleResult(err)
	}
	return events.MakeOKHandleResult()
}
