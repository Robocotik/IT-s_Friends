package postgres

import (
	"context"


	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"

	"fmt"
	"os"

	"github.com/mymmrac/telego"
)

func (psql Postgres) GetCathedras(bot *telego.Bot, chatID int64, identity *structures.Identity) []string {

	var res []string
	var cathedraTitle string
	rows, err := psql.Conn.Query(context.Background(), `
	SELECT DISTINCT ca.title
	FROM cathedras ca
	JOIN schedule s ON ca.id = s.cathedra_id
	JOIN fillials fi ON s.fillial_id = fi.id
	JOIN faculties fa ON s.faculty_id = fa.id
	WHERE fi.title = $1 AND fa.title = $2;
`, identity.Filial, identity.Faculty)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed in getting cathedras: %v\n", err)
		output.RiseError(bot, chatID, err)
		return []string{""}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cathedraTitle)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to scan cathedra title: %v\n", err)
			output.RiseError(bot, chatID, err)
			return []string{""}
		}
		res = append(res, cathedraTitle)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred during row iteration: %v\n", err)
		output.RiseError(bot, chatID, err)
		return []string{""}
	}

	fmt.Println("I FOUND CATHEDRAS: ", res)
	return res

}
