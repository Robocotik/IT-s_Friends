package logic

import (
	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func SearchGroupUID(bot *telego.Bot, msg telego.Message, conn *pgx.Conn, identity *structures.Identity) string {
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
		identity.Group, identity.Filial,
	).Scan(&res)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed in getting group uid: %v\n", err)
		output.RiseError(bot, msg, err)
		return ""
	}
	fmt.Println("UUID: ", res)

	return res
}
