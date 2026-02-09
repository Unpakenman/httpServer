package bootstrap

import (
	"fmt"
	"google.golang.org/grpc"
	"httpServer/internal/app/config"
	"log/slog"
	"net"
)

func NewGRPCPortListener(cfg *config.GRPCServerConfig,
) (net.Listener, error) {
	if cfg == nil {
		return nil, fmt.Errorf("invalid grpc config")
	}
	return net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
}

func NewGRPCServer(cfg *config.GRPCServerConfig,
	log *slog.Logger,
	interceptors ...grpc.ServerOption) (*grpc.Server, error) {
	if cfg == nil {
		return nil, fmt.Errorf("invalid grpc config")
	}
	grpcServer := grpc.NewServer(interceptors...)
	return grpcServer, nil
}

type GRPCResponse[T any, K any] struct {
	Data  *T
	Error *GRPCResponseError[K]
}
type GRPCResponseError[T any] struct {
	Code    int32
	Message string
	Details T
}
