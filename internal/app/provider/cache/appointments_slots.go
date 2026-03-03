package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"httpServer/internal/app/provider/models"
	"time"
)

func (c *redisCache) SetAppointmentsSlots(
	ctx context.Context,
	key string,
	data models.AppointmentsSlots,
	ttl time.Duration,
) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal appointments slot cache: %w", err)
	}
	if err := c.Set(ctx, key, dataBytes, ttl); err != nil {
		return fmt.Errorf("set appointments slot cache: %w", err)
	}
	return nil
}

func (c *redisCache) GetAppointmentsSlots(
	ctx context.Context,
	key string,
) (*models.AppointmentsSlots, error) {
	item, _ := c.Get(ctx, key)
	if len(item) == 0 {
		return nil, nil
	}
	var data models.AppointmentsSlots
	err := json.Unmarshal(item, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal appointments slot cache: %w", err)
	}
	return &data, nil
}
