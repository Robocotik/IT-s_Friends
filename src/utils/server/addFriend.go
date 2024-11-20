package server

import (
	"Friends/src/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	"os"
)

func AddFriend(bot *telego.Bot, msg telego.Message, conn *pgx.Conn, nickname string, group_uid string) (int64, error) {
	var id int64
	fmt.Printf("\n Я ДОБАВИЛ %s | %s : \n", nickname, group_uid)

	err := conn.QueryRow(
		context.Background(),
		"INSERT INTO friends (nickname, group_id) VALUES ($1, $2) ON CONFLICT (nickname, group_id) DO NOTHING  RETURNING (friend_id)",
		nickname, group_uid,
	).Scan(&id)
	fmt.Println("Я ВЕРНУЛ ID друга: ", id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed in add friends: %v\n", err)
		utils.RiseError(bot, msg, err)
		return 0, err
	}

	return id, nil // Возвращаем ID вставленной строки
}
