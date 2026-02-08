package provider

import (
	"fmt"
	pgclient "httpServer/internal/app/client/pg"
	"log"
	"reflect"
)

type goExampleDBProvider struct {
	conn pgclient.PGClient
}

func NewGoExampleDBProvider(dbConn pgclient.PGClient) GoExampleProvider {
	return &goExampleDBProvider{
		conn: dbConn,
	}
}

func (p *goExampleDBProvider) BeginTransaction() (pgclient.Transaction, error) {
	return p.conn.BeginTransaction()
}

func (p *goExampleDBProvider) RollbackTransaction(tx pgclient.Transaction) {
	if tx == nil || reflect.ValueOf(tx).IsNil() {
		return
	}
	err := tx.Rollback()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (p *goExampleDBProvider) CommitTransaction(tx pgclient.Transaction) error {
	if tx == nil || reflect.ValueOf(tx).IsNil() {
		return fmt.Errorf("nil transaction pointer in CommitTransaction")
	}
	return tx.Commit()
}
