package cache

import (
	"context"
	"httpServer/internal/app/provider/models"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error

	SetAppointmentsSlots(ctx context.Context, key string, data models.AppointmentsSlots, ttl time.Duration) error
	GetAppointmentsSlots(ctx context.Context, key string) (*models.AppointmentsSlots, error)
}
