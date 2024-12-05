package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewDatabase(DB_DRIVER string, DB_USER string, DB_PASSWORD string, DB_HOST string, DB_PORT string, DB_TABLE string) (*pgx.Conn, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s", os.Getenv(DB_DRIVER), os.Getenv(DB_USER), os.Getenv(DB_PASSWORD), os.Getenv(DB_HOST), os.Getenv(DB_PORT), os.Getenv(DB_TABLE))
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return conn, err
	}
	return conn, nil
}