package assets

import (
	"context"
	"fmt"
	"os"

	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func GetGroups(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, identity *structures.Identity) []string {

	var res []string
	var groupsTitle string

	rows, err := conn.Query(context.Background(), `
	SELECT g.title
	FROM groups g
	JOIN schedule s ON g.id = s.group_id
	JOIN cathedras c ON s.cathedra_id = c.id
	JOIN fillials fi ON s.fillial_id = fi.id
	WHERE c.title IN ($1) 
	  AND fi.title IN ($2);
`, identity.Cathedra, identity.Filial)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed in getting groups: %v\n", err)
		output.RiseError(bot, msg, err)
		return []string{""}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&groupsTitle)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to scan group title: %v\n", err)
			output.RiseError(bot, msg, err)
			return []string{""}
		}
		res = append(res, groupsTitle)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred during row iteration: %v\n", err)
		output.RiseError(bot, msg, err)
		return []string{""}
	}
	return res

}
