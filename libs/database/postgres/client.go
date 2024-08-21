package postgres

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	logger *slog.Logger
	config Config
	db     DB
}

func ProvideClient(ctx context.Context, config Config, log *slog.Logger) (*Client, error) {
	conn, err := pgxpool.New(ctx, config.DSN())

	if err != nil {
		return nil, err
	}

	return &Client{
		logger: log,
		config: config,
		db:     pg(conn, log),
	}, nil
}

func (c *Client) DB() DB {
	return c.db
}

func (c *Client) Ping(ctx context.Context) error {
	return c.db.Ping(ctx)
}

func (c *Client) Close() error {
	return c.db.Close()
}
