package bootstrap

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"httpServer/internal/app/config"
	"log/slog"
	"net/http"
	"time"
)

func NewChiRouter() *chi.Mux {
	return chi.NewRouter()
}

func RunHTTPServer(
	chiRouter *chi.Mux,
	logger slog.Logger,
	cfg *config.HTTPServerConfig,
) (*http.Server, error) {
	if cfg == nil {
		return nil, errors.New("invalid http config")
	}

	server := &http.Server{
		Addr:              fmt.Sprintf(`:%d`, cfg.ServerPort),
		Handler:           chiRouter,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(err.Error())
		}
	}()

	return server, nil
}
