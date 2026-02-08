package middleware

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func OtelhttpLabelerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		labeler := &otelhttp.Labeler{}
		labeler.Add(attribute.String(string(semconv.HTTPRouteKey), r.URL.Path))

		ctx := otelhttp.ContextWithLabeler(r.Context(), labeler)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
