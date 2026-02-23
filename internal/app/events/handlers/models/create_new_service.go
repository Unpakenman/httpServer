package models

import "github.com/shopspring/decimal"

type CreateNewServiceRequest struct {
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Price           decimal.Decimal `json:"price"`
	DurationMinutes int64           `json:"duration_minutes"`
	IsActive        bool            `json:"is_active"`
}
