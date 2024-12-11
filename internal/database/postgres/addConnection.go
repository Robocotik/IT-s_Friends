package postgres

import (
	"context"
	"errors"
	"fmt"
	"os"

	errorsCustom "github.com/Robocotik/IT-s_Friends/internal/models/errors"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mymmrac/telego"
)


func (psql Postgres) AddConnection(ctx context.Context, bot *telego.Bot, chatID int64, user_id int64, friend_id int64) error {
	// fmt.Printf("\n Я ДОБАВИЛ %s в %s : \n", user_id, friend_id)
	_, err := psql.Conn.Exec(
		ctx,
		"INSERT INTO user_friend (user_id, friend_id) VALUES ($1, $2)",
		user_id, friend_id,
	)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23503" { // Код ошибки для нарушения ограничения
				fmt.Fprintf(os.Stderr, "Insert failed: constraint violation on name_length: %v\n", pgErr.Message)
				output.WriteMessage(bot, chatID, errorsCustom.ErrFriendAlreadyAdded_23503)
				return errors.New(errorsCustom.ErrFriendAlreadyAdded_23503)
			}
		}
		fmt.Fprintf(os.Stderr, "QueryRow failed in connection : %v\n", err)
		output.RiseError(bot, chatID, err)
		return err
	}
	return nil
}
