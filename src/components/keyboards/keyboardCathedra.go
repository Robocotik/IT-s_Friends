package keyboard

import (
	"Friends/src/assets"
	"Friends/src/utils"
	// "Friends/src/messages"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	// "fmt"
)

func CreateKeyboardCathedra(filial string, faculty string) *telego.ReplyKeyboardMarkup {
	var cathedras = assets.GetCathedras(filial, faculty)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(cathedras)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите кафедру").WithOneTimeKeyboard()
}
