package database

import (
	// "Friends/src/utils"
	// "context"
	// "fmt"
	// "os"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func UpdateBd(bot *telego.Bot, msg telego.Message, conn *pgx.Conn) error {
	// fmt.Printf("\n Я ОБновил бд : \n")
	// request := logic.

	// _, err := conn.Exec(
	// 	context.Background(),
	// 	"INSERT INTO shedule (course_title, faculty_title, cathedra_title, group_title, timetable) VALUES ($1, $2, $3, $4, $5)",
	// 	course_title, course_title, course_title , course_title , course_title
	// )
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed in connection : %v\n", err)
	// 	utils.RiseError(bot, msg, err)
	// 	return err
	// }
	return nil
}
