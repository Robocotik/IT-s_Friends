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
		`
		SELECT uuid FROM schedule
		JOIN fillials ON schedule.fillial_id = fillials.id
		JOIN groups ON schedule.group_id = groups.id
		WHERE groups.title = $1 AND fillials.title = $2;
		 `,
		friend.Group, friend.Filial,
	).Scan(&res)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed in getting group uid: %v\n", err)
		utils.RiseError(bot, msg, err)
		return ""
	}
	fmt.Println("UUID: ", res)

	return res
}
