package httpserver

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"httpServer/internal/app/config"
	"httpServer/internal/app/httpserver/clinics"
	"httpServer/internal/app/httpserver/mapper"
	"httpServer/internal/app/httpserver/middleware"
	useCase "httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
	"log/slog"
)

type HttpServer interface {
}

type httpServer struct {
	clinics.HttpRouter
}

func NewHttpServer(
	logger slog.Logger,
	chiRouter *chi.Mux,
	httpConfig *config.HTTPServerConfig,
	mapperInstance mapper.Mapper,
	validatorInstance validator.Validator,
	clinicUseCase useCase.UseCase,
) HttpServer {
	chiRouter.Use(chimiddleware.Recoverer)
	apiGroup := chiRouter.Route(httpConfig.ApiDefaultPath, func(apiGroup chi.Router) {
		apiGroup.Use(middleware.OtelhttpLabelerMiddleware)
	})

	return &httpServer{
		clinics.NewHttpRouter(
			logger,
			apiGroup,
			httpConfig,
			mapperInstance,
			validatorInstance,
			clinicUseCase,
		),
	}
}
