package clinics

import (
	"github.com/go-chi/chi/v5"
	"httpServer/internal/app/config"
	"httpServer/internal/app/httpserver/mapper"
	"httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
	"log/slog"
)

type HttpRouter interface {
}

type httpRouter struct {
	logger    slog.Logger
	mapper    mapper.Mapper
	validator validator.Validator
	usecase   clinics.UseCase
}

func NewHttpRouter(
	logger slog.Logger,
	chiRouter chi.Router,
	httpConfig *config.HTTPServerConfig,
	mapper mapper.Mapper,
	validator validator.Validator,
	usecase clinics.UseCase,
) HttpRouter {
	router := &httpRouter{
		logger:    logger,
		mapper:    mapper,
		validator: validator,
		usecase:   usecase,
	}
	chiRouter.Get("/ping", router.Ping)
	chiRouter.Post("/create_patient", router.CreatePatient)
	return router
}
