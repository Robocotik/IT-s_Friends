package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"github.com/mymmrac/telego"
)

func (psql Postgres) UpdateUser(bot *telego.Bot, chatID int64, identity structures.Identity, id int64) error {

	_, err := psql.Conn.Exec(context.Background(), `
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
		output.RiseError(bot, chatID, err)
		return err
	}

	return nil
}
