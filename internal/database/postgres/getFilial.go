package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"github.com/mymmrac/telego"
)

func (psql Postgres) GetFilials(bot *telego.Bot, chatID int64) []string {
	var res []string
	var filialTitle string

	rows, err := psql.Conn.Query(
		context.Background(),
		"SELECT title FROM fillials",
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed in getting filials: %v\n", err)
		output.RiseError(bot, chatID, err)
		return []string{""}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&filialTitle)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to scan filial title: %v\n", err)
			output.RiseError(bot, chatID, err)
			return []string{""}
		}
		res = append(res, filialTitle)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred during row iteration: %v\n", err)
		output.RiseError(bot, chatID, err)
		return []string{""}
	}
	fmt.Println(res)
	return res
}
