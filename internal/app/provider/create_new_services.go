package provider

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	pgclient "httpServer/internal/app/client/pg"
)

type CreateNewServiceRequest struct {
	Name            string
	Description     string
	Price           decimal.Decimal
	DurationMinutes int64
	IsActive        bool
}

func (p *goExampleDBProvider) CreateNewService(
	ctx context.Context,
	tx pgclient.Transaction,
	data CreateNewServiceRequest,
) error {
	_, err := p.conn.Exec(
		ctx,
		"CreateNewService",
		nil,
		tx,
		data.Name,
		data.Description,
		data.Price,
		data.DurationMinutes,
		data.IsActive)
	if err != nil {
		return fmt.Errorf("could not create new service: %w", err)
	}
	return nil
}
