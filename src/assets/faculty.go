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

func GetFaculties(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, identity *structures.Identity) []string {

	var res []string
	var facultyTitle string
	rows, err := conn.Query(context.Background(), `
	SELECT DISTINCT f.title
	FROM faculties f
	JOIN schedule s ON f.id = s.faculty_id
	JOIN fillials fi ON s.fillial_id = fi.id
	WHERE fi.title = $1;
`, identity.Filial)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed in getting faculties: %v\n", err)
		utils.RiseError(bot, msg, err)
		return []string{""}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&facultyTitle)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to scan faculty title: %v\n", err)
			utils.RiseError(bot, msg, err)
			return []string{""}
		}
		res = append(res, facultyTitle)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred during row iteration: %v\n", err)
		utils.RiseError(bot, msg, err)
		return []string{""}
	}

	fmt.Println("I FOUND FACULTIES: ", res)
	return res
}
