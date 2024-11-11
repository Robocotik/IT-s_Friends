package keyboard

import (
	"Friends/src/assets"
	"Friends/src/utils"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardFilial() *telego.ReplyKeyboardMarkup {
	var filials = assets.GetFilials()
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(filials)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите филиал").WithOneTimeKeyboard()
}
