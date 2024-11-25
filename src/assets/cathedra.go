package assets

import (
	"github.com/Robocotik/IT-s_Friends/src/components/structures"
	"github.com/Robocotik/IT-s_Friends/src/utils"
	"context"

	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	"os"
)

func GetCathedras(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, identity *structures.Identity) []string {

	var res []string
	var cathedraTitle string
	rows, err := conn.Query(context.Background(), `
	SELECT DISTINCT ca.title
	FROM cathedras ca
	JOIN schedule s ON ca.id = s.cathedra_id
	JOIN fillials fi ON s.fillial_id = fi.id
	JOIN faculties fa ON s.faculty_id = fa.id
	WHERE fi.title = $1 AND fa.title = $2;
`, identity.Filial, identity.Faculty)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed in getting cathedras: %v\n", err)
		utils.RiseError(bot, msg, err)
		return []string{""}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cathedraTitle)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to scan cathedra title: %v\n", err)
			utils.RiseError(bot, msg, err)
			return []string{""}
		}
		res = append(res, cathedraTitle)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred during row iteration: %v\n", err)
		utils.RiseError(bot, msg, err)
		return []string{""}
	}

	fmt.Println("I FOUND CATHEDRAS: ", res)
	return res

}
