package models

import "github.com/shopspring/decimal"

type Services struct {
	ServiceID       int64           `db:"service_id"`
	Name            string          `db:"name"`
	Description     string          `db:"description"`
	Price           decimal.Decimal `db:"price"`
	DurationMinutes int64           `db:"duration_minutes"`
	IsActive        bool            `db:"is_active"`
}
