package server

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func GetFavsFromId(conn *pgx.Conn, id string) ([]string, error) {
	var favourites []string

	err := conn.QueryRow(
		context.Background(),
		"SELECT favourites FROM users WHERE id = $1",
		id,
	).Scan(&favourites)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return favourites, nil
}