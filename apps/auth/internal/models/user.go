package dto

import (
	"database/sql"
	"time"
)

type User struct {
	ID         string
	email      string
	username   string
	password   string
	created_at time.Time
	updated_at time.Time
}

type CreateUser struct {
	email    string
	username string
	password string
}

type UpdateUser struct {
	username sql.NullString
	password sql.NullString
}