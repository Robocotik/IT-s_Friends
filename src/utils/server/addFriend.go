package server

import (
	"Friends/src/components/structures"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mymmrac/telego"
)

func AddFriend(bot *telego.Bot, msg telego.Message, conn *pgx.Conn, friend *structures.AskedFriend) (int64, error) {
	var id int64

	// Сначала проверяем, существует ли запись
	err := conn.QueryRow(
		context.Background(),
		"SELECT friend_id FROM friends WHERE nickname = $1 AND group_id = $2",
		friend.NickName, friend.Uuid,
	).Scan(&id)

	if err != nil && err == pgx.ErrNoRows {
		err = conn.QueryRow(
			context.Background(),
			"INSERT INTO friends (nickname, group_id, group_title) VALUES ($1, $2, $3) RETURNING friend_id",
			friend.NickName, friend.Uuid, friend.Group,
		).Scan(&id)
		if err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok {
				if pgErr.Code == "23514" { // Код ошибки для нарушения ограничения
					fmt.Fprintf(os.Stderr, "Insert failed: constraint violation on name_length: %v\n", pgErr.Message)
					return -1, errors.New("too long")
				}
			}
			fmt.Fprintf(os.Stderr, "Insert failed: %v\n", err)
			return -1, err
		}
	} else {

		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}

	return id, nil
}
