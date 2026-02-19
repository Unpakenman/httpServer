package middleware

import (
	"bytes"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	logger "httpServer/internal/app/log"
	"io"
	"net/http"
	"time"
)

func CommonMiddleware(logger logger.LogClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		commonMiddlewareHandler := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Error(err)
			}

			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			ctx := r.Context()
			methodName := r.Method + " " + r.URL.Path

			logger.InfoCtx(ctx, "http_server.request "+methodName+", Body: "+string(bodyBytes))

			start := time.Now()

			next.ServeHTTP(ww, r)

			duration := time.Since(start)
			statusCode := ww.Status()

			requestInfoText := fmt.Sprintf(
				"http_server.response HTTP method: %s, time spent: %v, status code: %d",
				methodName,
				duration,
				statusCode,
			)

			if err != nil {
				logger.ErrorCtx(ctx, err)
			} else {
				logger.InfoCtx(ctx, requestInfoText)
			}
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			otelHandler := otelhttp.NewHandler(http.HandlerFunc(commonMiddlewareHandler), r.URL.Path)
			otelHandler.ServeHTTP(w, r)
		})
	}
}
