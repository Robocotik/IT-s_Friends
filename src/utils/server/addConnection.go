package server

import (
	"Friends/src/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func AddConnection(bot *telego.Bot, msg telego.Message, conn *pgx.Conn, user_id int64, friend_id int64) error {
	fmt.Printf("\n Я ДОБАВИЛ %s в %s : \n", user_id, friend_id)
	_, err := conn.Exec(
		context.Background(),
		"INSERT INTO user_friend (user_id, friend_id) VALUES ($1, $2)",
		user_id, friend_id,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed in connection : %v\n", err)
		utils.RiseError(bot, msg, err)
		return err
	}
	return nil
}
