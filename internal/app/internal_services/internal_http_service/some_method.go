package internal_http_service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"httpServer/internal/app/policy"
	"io"
	"net/http"
	"time"
)

func (s *internalServiceAPI) SomeExample(ctx context.Context, req ActionRequest) (*ActionResponse, error) {

	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	reqWithCtx, err := http.NewRequestWithContext(ctx, http.MethodPost, s.baseURL+ActionURL, bytes.NewBuffer(reqData))
	if err != nil {
		return nil, err
	}

	reqWithCtx.Header.Set("Content-Type", "application/json")

	respRaw, err := s.httpClient.Do(
		reqWithCtx,
		policy.WithHttpRetryPolicy(3, time.Millisecond*100, checkRetry),
		policy.WithTimeout[*http.Response](time.Second*10),
		policy.WithHttpCircuitBreaker(reqWithCtx.URL.Path, 0.8, time.Second*10, checkCircuitBreakerError),
	)
	if err != nil {
		return nil, err
	}

	if respRaw == nil {
		return nil, errors.New(`invalid response, nil response value`)
	}

	defer respRaw.Body.Close()

	if respRaw.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`invalid response, status code %d`, respRaw.StatusCode)
	}

	buf, err := io.ReadAll(respRaw.Body)
	if err != nil {
		return nil, err
	}

	response := &ActionResponse{}
	err = json.Unmarshal(buf, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
