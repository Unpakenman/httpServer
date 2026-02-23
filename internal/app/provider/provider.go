package provider

import (
	"context"
	pgclient "httpServer/internal/app/client/pg"
)

type goExampleDBProvider struct {
	conn pgclient.PGClient
}

func NewGoExampleDBProvider(dbConn pgclient.PGClient) GoExampleProvider {
	return &goExampleDBProvider{
		conn: dbConn,
	}
}

func (p *goExampleDBProvider) WithTransaction(ctx context.Context, fn func(context.Context, pgclient.Transaction) error) error {
	return p.conn.WithTransaction(ctx, fn)
}
