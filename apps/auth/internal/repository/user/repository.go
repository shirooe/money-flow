package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/shirooe/money-flow/apps/auth/internal/models"
	"github.com/shirooe/money-flow/libs/database/postgres"
)

func (ur *userRepository) Create(ctx context.Context, user models.) (string, error) {
	sql, args, err := sq.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("email", "username", "password").
		Values(user.Email, user.Username, user.Password).
		Suffix("RETURNING id").ToSql()

	if err != nil {
		return "", err
	}

	query := postgres.Query{
		Name:     "create_user",
		QueryRaw: sql,
	}

	var id string
	if err := ur.db.QueryRowContext(ctx, query, args...).Scan(id); err != nil {
		return "", err
	}

	return id, nil
}
