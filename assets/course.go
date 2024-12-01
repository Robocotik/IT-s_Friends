package assets

import (
	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func GetCourses(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, identity *structures.Identity) []string {

	var res []string
	var coursesTitle string

	rows, err := conn.Query(context.Background(), `
	SELECT DISTINCT c.title
	FROM courses c
	JOIN schedule s ON c.id = s.course_id
	JOIN fillials fi ON s.fillial_id = fi.id
	JOIN cathedras ca ON s.cathedra_id = ca.id
	WHERE fi.title = $1 AND ca.title = $2;
`, identity.Filial, identity.Cathedra)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed in getting courses: %v\n", err)
		output.RiseError(bot, msg, err)
		return []string{""}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&coursesTitle)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to scan course title: %v\n", err)
			output.RiseError(bot, msg, err)
			return []string{""}
		}
		res = append(res, coursesTitle)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred during row iteration: %v\n", err)
		output.RiseError(bot, msg, err)
		return []string{""}
	}

	fmt.Println("I FOUND COURSES: ", res)
	return res

}