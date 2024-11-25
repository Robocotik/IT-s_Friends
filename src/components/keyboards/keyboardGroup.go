package keyboard

import (
	"Friends/src/assets"
	"Friends/src/components/structures"
	"Friends/src/utils"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardGroup(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, identity *structures.Identity) *telego.ReplyKeyboardMarkup {
	var groups = assets.GetGroups(conn, bot, msg, identity)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(groups)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите группу").WithOneTimeKeyboard()
}
