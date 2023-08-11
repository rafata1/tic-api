package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

// Readonly for wrapping sqlx functionalities
type Readonly interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// Transaction for wrapping sqlx functionalities
type Transaction interface {
	Readonly

	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

var _ Transaction = &sqlx.DB{}
var _ Transaction = &sqlx.Tx{}

// Provider for creating Readonly and Transaction
//
//go:generate moq --out provider_mocks.go . Provider
type Provider interface {
	Transact(ctx context.Context, fn func(ctx context.Context) error) error
	Readonly(ctx context.Context) context.Context
}

type providerImpl struct {
	db *sqlx.DB
}

// NewProvider ...
func NewProvider(db *sqlx.DB) Provider {
	return &providerImpl{db: db}
}

// Transact ...
func (p *providerImpl) Transact(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	_, ok := getTxFromContext(ctx)
	if ok {
		return fn(ctx)
	}

	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			err = tx.Rollback()
			panic(r)
		} else if err != nil {
			_ = tx.Rollback()
		}
	}()

	ctx = context.WithValue(ctx, ctxTxKey, ctxTxValue{
		tx: tx,
	})

	err = fn(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// Readonly ...
func (p *providerImpl) Readonly(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxReadonlyKey, ctxReadonlyValue{
		db: p.db,
	})
}

func getTxFromContext(ctx context.Context) (ctxTxValue, bool) {
	tx, ok := ctx.Value(ctxTxKey).(ctxTxValue)
	return tx, ok
}

// GetTx get Transaction from context
func GetTx(ctx context.Context) Transaction {
	tx, ok := getTxFromContext(ctx)
	if !ok {
		panic("Not found transaction")
	}
	return tx.tx
}

// GetReadonly get Readonly from context
func GetReadonly(ctx context.Context) Readonly {
	tx, ok := getTxFromContext(ctx)
	if ok {
		return tx.tx
	}

	db, ok := ctx.Value(ctxReadonlyKey).(ctxReadonlyValue)
	if ok {
		return db.db
	}

	panic("Not found readonly repository")
}

type ctxTxKeyType struct {
}

type ctxReadonlyKeyType struct {
}

var ctxTxKey = ctxTxKeyType{}
var ctxReadonlyKey = ctxReadonlyKeyType{}

type ctxTxValue struct {
	tx *sqlx.Tx
}

type ctxReadonlyValue struct {
	db *sqlx.DB
}
