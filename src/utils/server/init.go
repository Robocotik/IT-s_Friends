package server

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitBD() *pgx.Conn {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/DB_USERS", os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
