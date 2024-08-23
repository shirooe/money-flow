package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Client struct {
	logger *zerolog.Logger
	config Config
	db     DB
}

func ProvideClient(ctx context.Context, config Config, log *zerolog.Logger) (*Client, error) {
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
