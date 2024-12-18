package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/internal/database"
	"github.com/Robocotik/IT-s_Friends/internal/services/utils"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardFilial(bd database.IBd, bot *telego.Bot, chatID int64) *telego.ReplyKeyboardMarkup {
	var filials = bd.GetFilials(bot, chatID)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(filials)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите филиал").WithOneTimeKeyboard()
}
