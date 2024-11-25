package assets

import (
	"github.com/Robocotik/IT-s_Friends/src/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func GetFilials(conn *pgx.Conn, bot *telego.Bot, msg telego.Message) []string {
	var res []string
	var filialTitle string

	rows, err := conn.Query(
		context.Background(),
		"SELECT title FROM fillials",
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed in getting filials: %v\n", err)
		utils.RiseError(bot, msg, err)
		return []string{""}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&filialTitle)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to scan filial title: %v\n", err)
			utils.RiseError(bot, msg, err)
			return []string{""}
		}
		res = append(res, filialTitle)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred during row iteration: %v\n", err)
		utils.RiseError(bot, msg, err)
		return []string{""}
	}

	fmt.Println("I FOUND FILIALS: ", res)
	return res
}
