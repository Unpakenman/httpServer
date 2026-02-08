package policy

import (
	"net/http"
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/circuitbreaker"
	"github.com/failsafe-go/failsafe-go/failsafehttp"
)

func WithHttpCircuitBreaker(
	key string,
	failureRateThreshold float64,
	delay time.Duration,
	handleIf checkHttpCircuitBreaker,
) failsafe.Policy[*http.Response] {
	mu.Lock()
	defer mu.Unlock()

	if br, ok := httpBreakers[key]; ok {
		return br
	}

	cb := circuitbreaker.NewBuilder[*http.Response]().
		HandleIf(handleIf).
		WithFailureRateThreshold(failureRateThreshold, FailureExecutionThreshold, FailureWindow).
		WithDelay(delay).
		Build()

	httpBreakers[key] = cb
	return cb
}

func WithHttpRetryPolicy(
	maxRetry int,
	delay time.Duration,
	checkRetryFunc checkHttpRetry,
) failsafe.Policy[*http.Response] {
	return failsafehttp.NewRetryPolicyBuilder().
		WithMaxRetries(maxRetry).
		WithDelay(delay).
		HandleIf(checkRetryFunc).
		Build()
}
