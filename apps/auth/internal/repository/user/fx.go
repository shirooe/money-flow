package user

import (
	"github.com/rs/zerolog"
	"github.com/shirooe/money-flow/libs/database/postgres"
)

type userRepository struct {
	db     postgres.DB
	logger *zerolog.Logger
}

func New(db postgres.DB, logger *zerolog.Logger) userRepository {
	return userRepository{
		db:     db,
		logger: logger,
	}
}
