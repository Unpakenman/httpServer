package http

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

type DefaultTransport struct {
	inner http.RoundTripper
	log   *slog.Logger
}

func NewTransport(
	inner http.RoundTripper,
	log *slog.Logger,
) *DefaultTransport {
	return &DefaultTransport{
		inner: inner,
		log:   log,
	}
}

func (t *DefaultTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	// Создание и идентификация запроса
	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID().String()
	methodName := req.Method + " " + req.URL.Path

	var requestID string
	if v := ctx.Value("request_id"); v != nil {
		if id, ok := v.(string); ok {
			requestID = id
		}
	}

	log := t.log.With(
		slog.String("trace_id", traceID),
		slog.String("request_id", requestID),
		slog.String("method", methodName),
	)

	// Чтение тела запроса
	var requestBody string
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			log.Error(
				"failed to read request body",
				slog.Any("error", err),
			)
		} else {
			requestBody = string(bodyBytes)
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	log.Info(
		"http client request",
		slog.String("body", requestBody),
	)

	start := time.Now()
	resp, err := t.inner.RoundTrip(req)

	statusCode := 500
	if resp != nil {
		statusCode = resp.StatusCode
	}
	duration := time.Since(start)

	if err != nil {
		log.Error(
			"http client request failed",
			slog.Any("error", errors.Join(
				err,
				fmt.Errorf(
					"method: %s, time spent: %v, status code: %d",
					methodName,
					duration,
					statusCode,
				),
			)),
		)
	}

	if resp != nil && resp.Body != nil {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Error("failed to read response body", slog.Any("error", err))
		}

		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		log.Info(
			"http client response",
			slog.String("method", methodName),
			slog.Duration("duration", duration),
			slog.Int("status_code", statusCode),
			slog.String("body", string(bodyBytes)),
		)
	}

	return resp, err
}
