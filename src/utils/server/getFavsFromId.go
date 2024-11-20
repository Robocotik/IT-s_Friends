package server

import (
	"Friends/src/components/structures"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetFavsFromId(conn *pgx.Conn, id int64) ([]structures.Fav, error) {
	favourites := []structures.Fav{}

	err := conn.QueryRow(
		context.Background(),
		"SELECT favourites::json FROM users WHERE id = $1",
		id,
	).Scan(&favourites)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		fmt.Fprintf(os.Stderr, "QueryRow failed in get friends: %v\n", err)
		return nil, err
	}

	return favourites, nil
}
