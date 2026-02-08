package internal_http_service

import (
	"context"
	"errors"
	"net/http"
	"slices"
)

func checkCircuitBreakerError(r *http.Response, err error) bool {
	if err != nil {
		return true
	}

	if r != nil && slices.Contains([]int{http.StatusInternalServerError, http.StatusServiceUnavailable, http.StatusGatewayTimeout}, r.StatusCode) {
		return true
	}

	return false
}

func checkRetry(r *http.Response, err error) bool {
	return err != nil && (errors.Is(err, context.DeadlineExceeded))
}
