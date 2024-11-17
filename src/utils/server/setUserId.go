package server

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func SetUserId(conn *pgx.Conn, id string) {
	fmt.Sprintln("Я ДОБАВИЛ ID1: ", id)
	_, err := conn.Exec(
		context.Background(),
		"insert into users (id) values ($1) on conflict (id) do nothing",
		id,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}
