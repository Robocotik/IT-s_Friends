package postgres

import (
	"context"
	"fmt"
	"os"


	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"github.com/mymmrac/telego"
)

func (psql Postgres) GetGroups(bot *telego.Bot, chatID int64, identity *structures.Identity) []string {

	var res []string
	var groupsTitle string

	rows, err := psql.Conn.Query(context.Background(), `
	SELECT g.title
	FROM groups g
	JOIN schedule s ON g.id = s.group_id
	JOIN courses c ON s.course_id = c.id
	JOIN fillials fi ON s.fillial_id = fi.id
	WHERE c.title = $1 
	  AND fi.title = $2;
`, identity.Course, identity.Filial)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed in getting groups: %v\n", err)
		output.RiseError(bot, chatID, err)
		return []string{""}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&groupsTitle)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to scan group title: %v\n", err)
			output.RiseError(bot, chatID, err)
			return []string{""}
		}
		res = append(res, groupsTitle)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred during row iteration: %v\n", err)
		output.RiseError(bot, chatID, err)
		return []string{""}
	}
	return res

}
