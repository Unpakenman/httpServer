package otelamqp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type MessageData struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

type consumerOtelProvider struct {
	tracer                   trace.Tracer
	messagingReceiveMessages metric.Int64Counter
	messagingProcessDuration metric.Float64Histogram
	messagingProcessMessages metric.Int64Counter

	config OtelConfig
}

type ConsumerOtelProvider interface {
	//Consumer
	StartConsume(
		ctx context.Context,
		routingKey string,
		message []byte,
	) context.Context
	ConsumeAck(ctx context.Context) error
	ConsumeReject(ctx context.Context, errorType string, err error) error
}

func NewConsumerOtelProvider(config OtelConfig) (ConsumerOtelProvider, error) {
	meter := otel.GetMeterProvider().Meter(
		scopeName,
		metric.WithInstrumentationVersion(scopeVersion),
		metric.WithSchemaURL(semconv.SchemaURL),
	)

	messagingReceiveMessages, err := meter.Int64Counter(
		semconv.MessagingReceiveMessagesName,
		metric.WithUnit(semconv.MessagingReceiveMessagesUnit),
		metric.WithDescription(semconv.MessagingReceiveMessagesDescription),
	)
	if err != nil {
		return nil, err
	}

	messagingProcessDuration, err := meter.Float64Histogram(
		semconv.MessagingProcessDurationName,
		metric.WithUnit(semconv.MessagingProcessDurationUnit),
		metric.WithDescription(semconv.MessagingProcessDurationDescription),
		metric.WithExplicitBucketBoundaries(defaultBounds...),
	)
	if err != nil {
		return nil, err
	}

	messagingProcessMessages, err := meter.Int64Counter(
		semconv.MessagingProcessMessagesName,
		metric.WithUnit(semconv.MessagingProcessMessagesUnit),
		metric.WithDescription(semconv.MessagingProcessMessagesDescription),
	)
	if err != nil {
		return nil, err
	}

	return &consumerOtelProvider{
		tracer: otel.GetTracerProvider().Tracer(
			scopeName,
			trace.WithInstrumentationVersion(scopeVersion),
			trace.WithSchemaURL(semconv.SchemaURL),
		),
		messagingReceiveMessages: messagingReceiveMessages,
		messagingProcessDuration: messagingProcessDuration,
		messagingProcessMessages: messagingProcessMessages,
		config:                   config,
	}, nil
}

func (o *consumerOtelProvider) StartConsume(
	ctx context.Context,
	routingKey string,
	messageBytes []byte,
) context.Context {

	var message MessageData
	err := json.Unmarshal(messageBytes, &message)
	if err != nil {
		message.UID = "empty-uid"
	}

	spanAttrs := []attribute.KeyValue{
		semconv.NetworkProtocolName(o.config.ProtocolName),
		semconv.MessagingSystemRabbitmq,
		semconv.MessagingDestinationName(o.config.QueueName),
		semconv.MessagingRabbitmqDestinationRoutingKey(routingKey),
		semconv.MessagingOperationTypeDeliver,
		semconv.ServerAddress(o.config.ServerName),
		semconv.ServerPort(o.config.ServerPort),
		semconv.MessagingClientID(o.config.ClientID),
		semconv.MessagingMessageBodySize(len(messageBytes)),
		semconv.MessagingMessageID(message.UID),
		semconv.MessagingDestinationTemporary(o.config.DestinationTemporary),
		semconv.MessagingDestinationAnonymous(o.config.DestinationAnonymous),
	}

	metricAttrs := []attribute.KeyValue{
		semconv.MessagingSystemRabbitmq,
		semconv.MessagingDestinationName(o.config.QueueName),
		semconv.MessagingRabbitmqDestinationRoutingKey(routingKey),
		semconv.MessagingOperationTypeDeliver,
		semconv.ServerAddress(o.config.ServerName),
		semconv.ServerPort(o.config.ServerPort),
	}

	ctx, _ = o.tracer.Start(
		ctx,
		fmt.Sprintf(`%s %s`, o.config.QueueName, semconv.MessagingOperationTypeDeliver.Value.AsString()),
		trace.WithSpanKind(trace.SpanKindConsumer),
		trace.WithAttributes(spanAttrs...),
	)

	ctx = context.WithValue(ctx, StartTimeKey, time.Now())
	ctx = context.WithValue(ctx, MetricAttrsKey, metricAttrs)

	o.messagingReceiveMessages.Add(ctx, 1, metric.WithAttributes(metricAttrs...))

	return ctx
}

func (o *consumerOtelProvider) ConsumeAck(ctx context.Context) error {
	span := trace.SpanFromContext(ctx)
	defer span.End()

	metricAttrs, ok := ctx.Value(MetricAttrsKey).([]attribute.KeyValue)
	if !ok {
		return errors.New(`can't get metric attributes from context`)
	}

	startTime, ok := ctx.Value(StartTimeKey).(time.Time)
	if !ok {
		return errors.New(`can't get startTime from context`)
	}

	span.SetStatus(codes.Ok, "Message successfully consumed")

	o.messagingProcessDuration.Record(ctx, time.Since(startTime).Seconds(), metric.WithAttributes(metricAttrs...))
	o.messagingProcessMessages.Add(ctx, 1, metric.WithAttributes(metricAttrs...))

	return nil
}

func (o *consumerOtelProvider) ConsumeReject(ctx context.Context, errorType string, err error) error {
	span := trace.SpanFromContext(ctx)
	defer span.End()

	metricAttrs, ok := ctx.Value(MetricAttrsKey).([]attribute.KeyValue)
	if !ok {
		return errors.New(`can't get metric attributes from context`)
	}

	startTime, ok := ctx.Value(StartTimeKey).(time.Time)
	if !ok {
		return errors.New(`can't get startTime from context`)
	}

	span.SetAttributes(semconv.ErrorTypeKey.String(errorType))
	metricAttrs = append(metricAttrs, semconv.ErrorTypeKey.String(errorType))

	span.SetStatus(codes.Error, fmt.Sprintf("Failed to consume message: %v", err))

	o.messagingProcessDuration.Record(ctx, time.Since(startTime).Seconds(), metric.WithAttributes(metricAttrs...))

	return nil
}
