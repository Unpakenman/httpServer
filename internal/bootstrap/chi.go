package bootstrap

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"httpServer/internal/app/config"
	"log"
	"net/http"
	"time"
)

func NewChiRouter() *chi.Mux {
	return chi.NewRouter()
}

func RunHTTPServer(
	chiRouter *chi.Mux,
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
			log.Fatal(err)
		}
	}()

	return server, nil
}
