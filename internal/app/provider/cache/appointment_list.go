package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"httpServer/internal/app/provider/models"
	"time"
)

func (c *redisCache) SetAppointmentsList(
	ctx context.Context,
	key string,
	data []models.AppointmentList,
	ttl time.Duration,
) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal appointments slot cache: %w", err)
	}
	if err := c.Set(ctx, key, dataBytes, ttl); err != nil {
		return fmt.Errorf("set appointments slot cache: %w", err)
	}
	fmt.Println("set appointments slot cache")
	return nil
}

func (c *redisCache) GetAppointmentsList(
	ctx context.Context,
	key string,
) (*[]models.AppointmentList, error) {
	item, _ := c.Get(ctx, key)
	if len(item) == 0 {
		return nil, nil
	}
	fmt.Println("GetAppointments from redis cache")
	var data []models.AppointmentList
	err := json.Unmarshal(item, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal appointments slot cache: %w", err)
	}
	return &data, nil
}
