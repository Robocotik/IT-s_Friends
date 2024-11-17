package server

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func AddIdToFavs(conn *pgx.Conn, user string, id_to_add string) {
	fmt.Printf("\n Я ДОБАВИЛ %s в %s : \n", id_to_add, user)
	_, err := conn.Exec(
		context.Background(),
		"UPDATE users SET favourites = array_append(favourites, $1) WHERE id = $2",
		id_to_add, user,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}
