package server

import (
	"Friends/src/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func AddUserId(bot *telego.Bot, msg telego.Message, conn *pgx.Conn, id int64, nickname string) error {
	fmt.Println("Я ДОБАВИЛ ID1: ", id, nickname)
	_, err := conn.Exec(context.Background(), "INSERT INTO users (id, nickname) VALUES ($1, $2) on conflict (id) do nothing", id, nickname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		utils.RiseError(bot, msg, err)
		return err
	}
	return nil
}