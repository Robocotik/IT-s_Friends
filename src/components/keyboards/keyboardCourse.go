package keyboard

import (
	"Friends/src/assets"
	"Friends/src/components/structures"
	"Friends/src/utils"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardCourse(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, friend *structures.AskedFriend) *telego.ReplyKeyboardMarkup {
	var courses = assets.GetCourses(conn, bot, msg, friend)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(courses)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите курс").WithOneTimeKeyboard()
}
