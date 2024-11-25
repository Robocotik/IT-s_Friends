package keyboard

import (
	"Friends/src/assets"
	"Friends/src/components/structures"
	"Friends/src/utils"

	// "Friends/src/messages"
	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	// "fmt"
)

func CreateKeyboardCathedra(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, friend *structures.AskedFriend) *telego.ReplyKeyboardMarkup {
	var cathedras = assets.GetCathedras(conn, bot, msg, friend)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(cathedras)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите кафедру").WithOneTimeKeyboard()
}
