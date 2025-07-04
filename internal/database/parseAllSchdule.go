package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Robocotik/IT-s_Friends/internal/models/entities"
	// "github.com/Robocotik/IT-s_Friends/internal/services/output"
	"io"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

type GroupPath struct {
	group_id    int64
	cathedra_id int64
	faculty_id  int64
	fillial_id  int64
	course_id   int64
}

func doRequest(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		fmt.Println("Ошибка при get запросе :", err)
		return []byte{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func pasteIntoSchedule(conn *pgx.Conn, group_uuid string, group_id, cathedra_id, faculty_id, fillial_id, course_id int64) error {
	_, err := conn.Exec(context.Background(),
		"INSERT INTO schedule (uuid, course_id, fillial_id, faculty_id, cathedra_id, group_id) VALUES($1, $2, $3, $4, $5, $6)",
		group_uuid, course_id, fillial_id, faculty_id, cathedra_id, group_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		// output.RiseError(bot, chatID, err)
		return err
	}
	return err
}

func parseCourse(conn *pgx.Conn, data entities.Course, fullpath *GroupPath) error {

	for _, group := range data.Children {
		err := conn.QueryRow(context.Background(),
			`
        INSERT INTO groups (title) VALUES($1)
        RETURNING id
    `,
			group.Abbr).Scan(&fullpath.group_id)
		if err != nil {
			// output.RiseError(bot, chatID, err)
			return err
		}

		pasteIntoSchedule(conn, group.Uuid, fullpath.group_id, fullpath.cathedra_id, fullpath.faculty_id, fullpath.fillial_id, fullpath.course_id)

	}
	return nil
}

func parseCathedra(conn *pgx.Conn, data entities.Cathedra, fullpath *GroupPath) error {

	for _, course := range data.Children {
		err := conn.QueryRow(context.Background(),
			`
        INSERT INTO courses (title) VALUES($1)
        RETURNING id
    `,
			course.Abbr).Scan(&fullpath.course_id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed in inserting group : %v\n", err)
			return err
		}
	parseCourse(conn, course, fullpath)
		// output.RiseError(bot, chatID, err)

	}
	return nil
}

func parseFaculty(conn *pgx.Conn, data entities.Faculty, fullpath *GroupPath) error {

	for _, cathedra := range data.Children {
		err := conn.QueryRow(context.Background(),
			`
        INSERT INTO cathedras (title) VALUES($1)
        RETURNING id
    
    `,
			cathedra.Abbr).Scan(&fullpath.cathedra_id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed in inserting cathedras: %v\n", err)
			return err
		}
		parseCathedra(conn, cathedra, fullpath)

		// output.RiseError(bot, chatID, err)

	}
	return nil
}

func parseFillial(conn *pgx.Conn, data entities.Fillial, fullpath *GroupPath) error {
	var err error

	for _, faculty := range data.Children {
		err = conn.QueryRow(context.Background(),
			"INSERT INTO faculties (title) VALUES($1) RETURNING id",
			faculty.Abbr).Scan(&fullpath.faculty_id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed in inserting faculties: %v\n", err)
			return err
		}
		parseFaculty(conn, faculty, fullpath)
	}
	return nil
}

func parseMainLayer(conn *pgx.Conn, result entities.Final) error {
	var fullpath GroupPath
	var err error

	for _, fillial := range result.Data.Children {
		err = conn.QueryRow(context.Background(),
			"INSERT INTO fillials (title) VALUES($1) RETURNING id",
			fillial.Abbr).Scan(&fullpath.fillial_id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed in inserting fillials: %v\n", err)
			return err
		}
		parseFillial(conn, fillial, &fullpath)
	}
	return nil
}

func ParseAllSchdule(ctx context.Context, conn *pgx.Conn, bot *telego.Bot, chatID int64) error {
	data, err := doRequest(ctx, "https://lks.bmstu.ru/lks-back/api/v1/structure")
	if err != nil {
		fmt.Println("Ошибка при чтении DATA:", err)
		return err
	}
	var result entities.Final
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return err
	}

	err = parseMainLayer(conn, result)
	if err != nil {
		fmt.Println("Ошибка при парсинге:", err)
		return err
	}
	return nil
}
