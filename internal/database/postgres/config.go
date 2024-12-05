package postgres

import (
	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	Conn        *pgx.Conn
	DB_DRIVER   string
	DB_USER     string
	DB_PORT     string
	DB_PASSWORD string
	DB_HOST     string
	DB_TABLE    string
}
