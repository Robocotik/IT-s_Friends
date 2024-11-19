package server

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func SetUserId(conn *pgx.Conn, id string, nickname string) error {
	fmt.Sprintln("Я ДОБАВИЛ ID1: %s | %s ", id, nickname)
	_, err := conn.Exec(
		context.Background(),
		"insert into users (id, nickname) values ($1, $2) on conflict (id) do nothing",
		id,
		nickname,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}
	return nil
}
