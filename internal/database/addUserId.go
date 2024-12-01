package database

import (
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func AddUserId(bot *telego.Bot, msg telego.Message, conn *pgx.Conn, id int64, nickname string) error {
	// fmt.Println("Я НАЧАЛ ДОБАВЛЯТЬ ID1: ", id, nickname)
	_, err := conn.Exec(context.Background(), "INSERT INTO users (id, nickname) VALUES ($1, $2) on conflict (id) do nothing", id, nickname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		output.RiseError(bot, msg, err)
		return err
	}
	// fmt.Println("Я ЗАКОНЧ ДОБАВЛЯТЬ ID1: ", id, nickname)
	return nil
}
