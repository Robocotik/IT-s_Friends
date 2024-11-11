package keyboard

import (
	"Friends/src/assets"
	"Friends/src/utils"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardGroup(filial string, faculty string, course string, cathedra string) *telego.ReplyKeyboardMarkup {
	var groups = assets.GetGroups(filial, faculty, course, cathedra)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(groups)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите группу").WithOneTimeKeyboard()
}
