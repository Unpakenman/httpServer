package http

import (
	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/failsafehttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"httpServer/internal/app/config"
	"log/slog"
	"net/http"
)

type Client interface {
	Do(
		req *http.Request,
		policies ...failsafe.Policy[*http.Response],
	) (*http.Response, error)
}

type client struct {
	baseClient *http.Client
}

func NewHTTPClient(
	cfg *config.HTTPClientConfig,
	log *slog.Logger,
) Client {

	return &client{
		baseClient: &http.Client{
			Timeout: cfg.Timeout,
			Transport: otelhttp.NewTransport(
				NewTransport(http.DefaultTransport, log),
				otelhttp.WithMetricAttributesFn(func(r *http.Request) []attribute.KeyValue {
					return []attribute.KeyValue{
						attribute.String(string(semconv.HTTPRouteKey), r.URL.Path),
					}
				}),
				otelhttp.WithSpanNameFormatter(
					func(operation string, r *http.Request) string {
						endpointName := r.URL.Path
						return operation + " " + endpointName
					},
				),
			),
		},
	}
}

func (c *client) Do(req *http.Request, policies ...failsafe.Policy[*http.Response]) (*http.Response, error) {
	return failsafehttp.NewRequest(req, c.baseClient, policies...).Do()
}
