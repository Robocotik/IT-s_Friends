package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/src/assets"
	"github.com/Robocotik/IT-s_Friends/src/components/structures"
	"github.com/Robocotik/IT-s_Friends/src/utils"

	// "Friends/src/messages"
	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	// "fmt"
)

func CreateKeyboardFaculty(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, identity *structures.Identity) *telego.ReplyKeyboardMarkup {
	var faculties = assets.GetFaculties(conn, bot, msg, identity)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(faculties)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите факультет").WithOneTimeKeyboard()
}
