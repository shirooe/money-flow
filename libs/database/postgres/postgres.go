package postgres

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

var _ DB = (*postgres)(nil)

type postgres struct {
	pool   *pgxpool.Pool
	logger *zerolog.Logger
}

func pg(conn *pgxpool.Pool, logger *zerolog.Logger) *postgres {
	return &postgres{
		pool:   conn,
		logger: logger,
	}
}

func (p *postgres) Pool() *pgxpool.Pool {
	return p.pool
}

func (p *postgres) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

func (p *postgres) Close() error {
	p.pool.Close()
	return nil
}

func (p *postgres) ScanOneContext(ctx context.Context, dest interface{}, query Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, query, args...)
	if err != nil {
		p.logger.Error().Msgf("failed to query %v", err)
		return err
	}
	return pgxscan.ScanOne(dest, rows)
}

func (p *postgres) ScanAllContext(ctx context.Context, dest interface{}, query Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, query, args...)
	if err != nil {
		p.logger.Error().Msgf("failed to query %v", err)
		return err
	}
	return pgxscan.ScanAll(dest, rows)
}

func (p *postgres) ScanRowContext(ctx context.Context, dest interface{}, query Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, query, args...)
	if err != nil {
		p.logger.Error().Msgf("failed to query %v", err)
		return err
	}
	return pgxscan.ScanRow(dest, rows)
}

func (p *postgres) ExecContext(ctx context.Context, query Query, args ...interface{}) (pgconn.CommandTag, error) {
	return p.pool.Exec(ctx, query.QueryRaw, args)
}

func (p *postgres) QueryContext(ctx context.Context, query Query, args ...interface{}) (pgx.Rows, error) {
	return p.pool.Query(ctx, query.QueryRaw, args)
}

func (p *postgres) QueryRowContext(ctx context.Context, query Query, args ...interface{}) pgx.Row {
	return p.pool.QueryRow(ctx, query.QueryRaw, args)
}
