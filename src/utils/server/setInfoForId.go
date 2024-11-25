package server

import (
	"github.com/Robocotik/IT-s_Friends/src/components/structures"
	"github.com/Robocotik/IT-s_Friends/src/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func SetInfoForId(bot *telego.Bot, msg telego.Message, conn *pgx.Conn, identity structures.Identity, id int64) error {

	_, err := conn.Exec(context.Background(), "INSERT INTO users (fillial, faculty, course, group, uuid) VALUES ($1, $2, $3, $4, $5) WHERE id = $6",
		identity.Filial, identity.Faculty, identity.Course, identity.Group, identity.Uuid, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed on setting info for id: %v\n", err)
		utils.RiseError(bot, msg, err)
		return err
	}

	return nil
}
