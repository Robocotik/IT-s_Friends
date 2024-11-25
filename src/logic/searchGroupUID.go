package logic

import (
	"github.com/Robocotik/IT-s_Friends/src/components/structures"
	"github.com/Robocotik/IT-s_Friends/src/utils"
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
		friend.Identity.Group, friend.Identity.Filial,
	).Scan(&res)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed in getting group uid: %v\n", err)
		utils.RiseError(bot, msg, err)
		return ""
	}
	fmt.Println("UUID: ", res)

	return res
}
