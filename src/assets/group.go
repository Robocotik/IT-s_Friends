package assets

import (
	"Friends/src/components/structures"
	"Friends/src/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func GetGroups(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, friend *structures.AskedFriend) []string {

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
`, friend.Cathedra, friend.Filial)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed in getting groups: %v\n", err)
		utils.RiseError(bot, msg, err)
		return []string{""}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&groupsTitle)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to scan group title: %v\n", err)
			utils.RiseError(bot, msg, err)
			return []string{""}
		}
		res = append(res, groupsTitle)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred during row iteration: %v\n", err)
		utils.RiseError(bot, msg, err)
		return []string{""}
	}

	fmt.Println("I FOUND FILIALS: ", res)
	return res

}
