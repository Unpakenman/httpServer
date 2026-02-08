package policy

import (
	"net/http"
	"sync"
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/timeout"
)

const (
	FailureExecutionThreshold = 10          // минимальное количество запросов после которых пойдет подсчет неудачных вызовов
	FailureWindow             = time.Minute // врменное окно в промежутке которого будут считаться каунтеры брейкера
)

var (
	mu           sync.Mutex
	httpBreakers = make(map[string]failsafe.Policy[*http.Response])
	grpcBreakers = make(map[string]failsafe.Policy[any])
)

type checkHttpRetry func(r *http.Response, err error) bool
type checkGrpcRetry func(r any, err error) bool

type checkHttpCircuitBreaker func(r *http.Response, err error) bool
type checkGrpcCircuitBreaker func(any, error) bool

func WithTimeout[T any](value time.Duration) failsafe.Policy[T] {
	return timeout.New[T](value)
}
