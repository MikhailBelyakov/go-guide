package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (
		int64,
		error,
	)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Ping(ctx context.Context) error
}

type segregatePool struct {
	wdb *pgxpool.Pool
	rdb *pgxpool.Pool
}

func New(wdb, rdb *pgxpool.Pool) Pool {
	return &segregatePool{
		wdb: wdb,
		rdb: rdb,
	}
}

func (sp *segregatePool) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return sp.rdb.Query(ctx, sql, args...)
}

func (sp *segregatePool) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return sp.rdb.QueryRow(ctx, sql, args...)
}

func (sp *segregatePool) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return sp.wdb.Exec(ctx, sql, arguments...)
}

func (sp *segregatePool) CopyFrom(
	ctx context.Context,
	tableName pgx.Identifier,
	columnNames []string,
	rowSrc pgx.CopyFromSource,
) (int64, error) {
	return sp.wdb.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

func (sp *segregatePool) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return sp.wdb.BeginTx(ctx, txOptions)
}

func (sp *segregatePool) Ping(ctx context.Context) error {
	if err := sp.wdb.Ping(ctx); err != nil {
		return err
	}

	if err := sp.rdb.Ping(ctx); err != nil {
		return err
	}

	return nil
}
