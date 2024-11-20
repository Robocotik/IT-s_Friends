package server

import (
	"Friends/src/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func AddIdToFavs(bot *telego.Bot, msg telego.Message, conn *pgx.Conn, user int64, id_to_add string) error {
	fmt.Printf("\n Я ДОБАВИЛ %s в %s : \n", id_to_add, user)
	_, err := conn.Exec(
		context.Background(),
		"UPDATE users SET favourites = array_append(favourites, $1) WHERE id = $2",
		id_to_add, user,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		utils.RiseError(bot, msg, err)
		return err
	}
	return nil
}
