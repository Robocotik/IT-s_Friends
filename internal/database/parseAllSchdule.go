package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Robocotik/IT-s_Friends/internal/models/entities"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func ParseAllSchdule(ctx context.Context, conn *pgx.Conn, bot *telego.Bot, msg telego.Message) error {
	resp, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://lks.bmstu.ru/lks-back/api/v1/structure", nil)
	var group_id, cathedra_id, faculty_id, fillial_id, course_id int64

	if err != nil {
		fmt.Println("Ошибка при get запросе :", err)
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении DATA:", err)
		return err
	}
	var result entities.Final
	err = json.Unmarshal(data, &result)
	for _, filial := range result.Data.Children {
		err := conn.QueryRow(context.Background(),
			`
    WITH ins AS (
        INSERT INTO fillials (title) VALUES($1)
        ON CONFLICT (title) DO NOTHING
        RETURNING id
    )
    SELECT COALESCE((SELECT id FROM ins), (SELECT id FROM fillials WHERE title = $1));
    `,
			filial.Abbr).Scan(&fillial_id)
		if err != nil {
			output.RiseError(bot, msg, err)
		}

		for _, faculty := range filial.Children {
			err := conn.QueryRow(context.Background(),
				`
    WITH ins AS (
        INSERT INTO faculties (title) VALUES($1)
        ON CONFLICT (title) DO NOTHING
        RETURNING id
    )
    SELECT COALESCE((SELECT id FROM ins), (SELECT id FROM faculties WHERE title = $1));
    `,
				faculty.Abbr).Scan(&faculty_id)
			if err != nil {
				output.RiseError(bot, msg, err)
			}
			for _, cathedra := range faculty.Children {
				err := conn.QueryRow(context.Background(),
					`
    WITH ins AS (
        INSERT INTO cathedras (title) VALUES($1)
        ON CONFLICT (title) DO NOTHING
        RETURNING id
    )
    SELECT COALESCE((SELECT id FROM ins), (SELECT id FROM cathedras WHERE title = $1));
    `,
					cathedra.Abbr).Scan(&cathedra_id)
				if err != nil {
					output.RiseError(bot, msg, err)
				}
				for _, course := range cathedra.Children {
					err := conn.QueryRow(context.Background(),
						`
    WITH ins AS (
        INSERT INTO courses (title) VALUES($1)
        ON CONFLICT (title) DO NOTHING
        RETURNING id
    )
    SELECT COALESCE((SELECT id FROM ins), (SELECT id FROM courses WHERE title = $1));
    `,
						course.Abbr).Scan(&course_id)
					if err != nil {
						output.RiseError(bot, msg, err)
					}
					for _, group := range course.Children {
						err := conn.QueryRow(context.Background(),
							`
    WITH ins AS (
        INSERT INTO groups (title) VALUES($1)
        ON CONFLICT (title) DO NOTHING
        RETURNING id
    )
    SELECT COALESCE((SELECT id FROM ins), (SELECT id FROM groups WHERE title = $1));
    `,
							group.Abbr).Scan(&group_id)
						if err != nil {
							output.RiseError(bot, msg, err)
						}
						_, err = conn.Exec(context.Background(),
							"INSERT INTO schedule (uuid, course_id, fillial_id, faculty_id, cathedra_id, group_id) VALUES($1, $2, $3, $4, $5, $6) on conflict (uuid) do nothing",
							group.Uuid, course_id, fillial_id, faculty_id, cathedra_id, group_id)
						if err != nil {
							fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
							output.RiseError(bot, msg, err)
							return err
						}

					}
				}
			}
		}
	}
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return err
	}
	return nil
}
