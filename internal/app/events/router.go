package events

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"go.opentelemetry.io/otel/trace"
	localerrors "httpServer/internal/app/errors"
	logger "httpServer/internal/app/log"
	"httpServer/internal/app/log/options"
	"httpServer/internal/app/rabbitmq_service/otelamqp"
	"strconv"
	"time"
)

type HandlersProcessor struct {
	handlers         map[string]Handler
	logger           logger.LogClient
	consumerProvider otelamqp.ConsumerOtelProvider
}

type Handler interface {
	Description() HandlerDescription
	Handle(ctx context.Context, data []byte, meta Meta) HandleResult
}

type HandleResult struct {
	Error error
	Code  string
}

type HandlerDescription struct {
	Alias       string
	Description string
}

type EventData[T any] struct {
	UID       string `json:"uid"`
	Name      string `json:"name"`
	EventID   int64  `json:"event_id"`
	Data      T      `json:"data"`
	CreatedAt string `json:"create_dttm"`
}

type Meta struct {
	TraceID   string `json:"trace_id"`
	EventUID  string `json:"event_uid"`
	RequestID string `json:"request_id"`
}

func NewRouter(
	handlers []Handler,
	logger logger.LogClient,
	consumerProvider otelamqp.ConsumerOtelProvider,
) *HandlersProcessor {
	processor := &HandlersProcessor{
		logger:           logger,
		handlers:         make(map[string]Handler),
		consumerProvider: consumerProvider,
	}
	for _, handler := range handlers {
		if handler == nil {
			logger.Warn("invalid handler passed to processor")
			continue
		}
		desc := handler.Description()
		processor.handlers[desc.Alias] = handler
	}
	return processor
}

func (h *HandlersProcessor) Process(routingKey string, data []byte) (err error) {
	ctx := h.consumerProvider.StartConsume(context.Background(), routingKey, data)

	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID().String()

	requestID := uuid.New()
	ctx = h.logger.SetOptionsToCtx(
		ctx,
		options.WithRequestID(requestID.String()),
		options.WithProtocol(options.AMQPProtocol),
		options.WithTraceID(traceID),
	)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
			h.logger.ErrorCtx(ctx, err)

			if consErr := h.consumerProvider.ConsumeReject(ctx, otelamqp.ErrorTypeInternalError, err); consErr != nil {
				h.logger.ErrorCtx(ctx, errors.Join(consErr, errors.New("error while sending reject data about metrics and traces to otelamqp collector")))
			}
		}
	}()

	eventData := &EventData[any]{}
	err = jsoniter.Unmarshal(data, eventData)
	if err != nil {
		h.logger.ErrorCtx(ctx, errors.Join(err, errors.New("unmarshal event data error")))

		if consErr := h.consumerProvider.ConsumeReject(ctx, otelamqp.ErrorTypeDecodeError, err); consErr != nil {
			h.logger.ErrorCtx(ctx, errors.Join(consErr, errors.New("error while sending reject data about metrics and traces to otelamqp collector")))
		}

		return err
	}

	h.logger.InfoCtx(ctx, "amqp_commands.request", eventData)

	handler, ok := h.handlers[eventData.Name]
	if !ok {
		err = fmt.Errorf("cannot find handler for event %s", eventData.Name)
		h.logger.ErrorCtx(ctx, err)

		if consErr := h.consumerProvider.ConsumeReject(ctx, otelamqp.ErrorTypeInternalError, err); consErr != nil {
			h.logger.ErrorCtx(ctx, errors.Join(consErr, errors.New("error while sending reject data about metrics and traces to otelamqp collector")))
		}

		return err
	}
	if handler == nil {
		err = fmt.Errorf("invalid handler for event %s", eventData.Name)
		h.logger.ErrorCtx(ctx, err)

		if consErr := h.consumerProvider.ConsumeReject(ctx, otelamqp.ErrorTypeInternalError, err); consErr != nil {
			h.logger.ErrorCtx(ctx, errors.Join(consErr, errors.New("error while sending reject data about metrics and traces to otelamqp collector")))
		}

		return err
	}

	meta := Meta{
		TraceID:   traceID,
		EventUID:  eventData.UID,
		RequestID: requestID.String(),
	}

	start := time.Now()
	result := handler.Handle(ctx, data, meta)
	duration := time.Since(start)

	h.logger.InfoCtx(ctx, fmt.Sprintf("amqp_commands.response time spent: %s", duration), eventData)
	if result.Error != nil {
		err = errors.Join(result.Error, fmt.Errorf("amqp error in event: %v, event uid: %v", eventData.Name, eventData.UID))
		h.logger.ErrorCtx(ctx, err)

		if consErr := h.consumerProvider.ConsumeReject(ctx, otelamqp.ErrorTypeInternalError, err); consErr != nil {
			h.logger.ErrorCtx(ctx, errors.Join(consErr, errors.New("error while sending reject data about metrics and traces to otelamqp collector")))
		}

		return err
	}

	err = h.consumerProvider.ConsumeAck(ctx)
	if err != nil {
		h.logger.ErrorCtx(ctx, errors.Join(err, errors.New("error while sending ack data about metrics and traces to otelamqp collector")))
	}

	return result.Error
}

func MakeInternalErrorHandleResult(err error) HandleResult {
	statusCode := localerrors.GetCodesByStatusName(localerrors.StatusInternalServerError)
	return HandleResult{
		Error: err,
		Code:  strconv.FormatUint(uint64(statusCode.HTTP), 10),
	}
}

func MakeOKHandleResult() HandleResult {
	statusCode := localerrors.GetCodesByStatusName(localerrors.StatusOK)
	return HandleResult{
		Error: nil,
		Code:  strconv.FormatUint(uint64(statusCode.HTTP), 10),
	}
}
