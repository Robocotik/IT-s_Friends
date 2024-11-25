package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/src/assets"
	"github.com/Robocotik/IT-s_Friends/src/utils"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardFilial(conn *pgx.Conn, bot *telego.Bot, msg telego.Message) *telego.ReplyKeyboardMarkup {
	var filials = assets.GetFilials(conn, bot, msg)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(filials)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите филиал").WithOneTimeKeyboard()
}
