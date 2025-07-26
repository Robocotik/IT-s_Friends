package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"github.com/mymmrac/telego"
)

func (psql Postgres) AddUserId(bot *telego.Bot, chatID int64, id int64, nickname string) error {
	// fmt.Println("Я НАЧАЛ ДОБАВЛЯТЬ ID1: ", id, nickname)
	_, err := psql.Conn.Exec(context.Background(), "INSERT INTO users (id, nickname) VALUES ($1, $2) on conflict (id) do nothing", id, nickname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed in adding user: %v\n", err)
		output.RiseError(bot, chatID, err)
		return err
	}
	// fmt.Println("Я ЗАКОНЧ ДОБАВЛЯТЬ ID1: ", id, nickname)
	return nil
}
