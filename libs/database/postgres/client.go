package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Client struct {
	logger zap.Logger
	config Config
	db     DB
}

func ProvideClient(ctx context.Context, config Config, logger *zap.Logger) (*Client, error) {
	conn, err := pgxpool.New(ctx, config.DSN())

	if err != nil {
		return nil, err
	}

	return &Client{
		logger: *logger,
		config: config,
		db:     pg(conn, logger),
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
