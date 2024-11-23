package logic

import (
	"Friends/src/components/structures"
	"Friends/src/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func SearchGroupUID(bot *telego.Bot, msg telego.Message, conn *pgx.Conn, friend *structures.AskedFriend) string {
	var res string

	// Исправлено: добавлена корректная SQL команда
	err := conn.QueryRow(
		context.Background(),
		"SELECT uuid FROM schedule WHERE group_title = $1 AND filial_title = $2",
		friend.Group, friend.Filial,
	).Scan(&res)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed in connection: %v\n", err)
		utils.RiseError(bot, msg, err)
		return ""
	}
	fmt.Println("UUID: ", res)

	return res
}
