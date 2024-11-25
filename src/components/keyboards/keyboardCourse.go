package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/src/assets"
	"github.com/Robocotik/IT-s_Friends/src/components/structures"
	"github.com/Robocotik/IT-s_Friends/src/utils"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardCourse(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, identity *structures.Identity) *telego.ReplyKeyboardMarkup {
	var courses = assets.GetCourses(conn, bot, msg, identity)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(courses)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите курс").WithOneTimeKeyboard()
}
