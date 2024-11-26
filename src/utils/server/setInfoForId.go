package server

import (
	"context"
	"fmt"
	"github.com/Robocotik/IT-s_Friends/src/components/structures"
	"github.com/Robocotik/IT-s_Friends/src/utils"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func SetInfoForId(bot *telego.Bot, msg telego.Message, conn *pgx.Conn, identity structures.Identity, id int64) error {

	_, err := conn.Exec(context.Background(), `
	UPDATE users SET fillial_title = $1,
	 faculty_title = $2, 
	 course_title = $3,
	 cathedra_title = $4,
	 group_title = $5, 
	 uuid = $6
	 WHERE id = $7;`,
		identity.Filial, identity.Faculty, identity.Course, identity.Cathedra, identity.Group, identity.Uuid, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed on setting info for id: %v\n", err)
		utils.RiseError(bot, msg, err)
		return err
	}

	return nil
}
