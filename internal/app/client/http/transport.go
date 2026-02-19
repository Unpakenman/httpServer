package http

import (
	"bytes"
	"errors"
	"fmt"
	logger "httpServer/internal/app/log"
	"io"
	"net/http"
	"time"
)

type DefaultTransport struct {
	inner  http.RoundTripper
	logger logger.LogClient
}

func NewTransport(
	inner http.RoundTripper,
	log logger.LogClient,
) *DefaultTransport {
	return &DefaultTransport{
		inner:  inner,
		logger: log,
	}
}

func (t *DefaultTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	// Создание и идентификация запроса
	methodName := req.Method + " " + req.URL.Path

	// Чтение тела запроса
	var requestBody string
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			t.logger.ErrorCtx(ctx, fmt.Errorf("failed to read request body: %w", err))
		} else {
			requestBody = string(bodyBytes)
			// Восстановление тела запроса для последующего использования
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	t.logger.InfoCtx(ctx, fmt.Sprintf(
		"method: %s, request body: %s",
		methodName,
		requestBody,
	))

	start := time.Now()
	resp, err := t.inner.RoundTrip(req)

	statusCode := 500
	if resp != nil {
		statusCode = resp.StatusCode
	}
	duration := time.Since(start)

	if err != nil {
		t.logger.ErrorCtx(
			ctx,
			errors.Join(err, fmt.Errorf("http_client.response %s, time spent: %v, status code: %d",
				methodName,
				duration,
				statusCode,
			)),
		)
	}

	if resp != nil && resp.Body != nil {

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.logger.ErrorCtx(ctx, err)
		}

		// Восстановление тела запроса
		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		t.logger.InfoCtx(ctx, fmt.Sprintf(
			"http_client.response %s, time spent: %v, status code: %d, body: %s",
			methodName,
			duration,
			statusCode,
			string(bodyBytes),
		))
	}

	return resp, err
}
