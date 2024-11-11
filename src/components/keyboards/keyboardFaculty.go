package keyboard

import (
	"Friends/src/assets"
	"Friends/src/utils"
	// "Friends/src/messages"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	// "fmt"
)

func CreateKeyboardFaculty(filial string) *telego.ReplyKeyboardMarkup {
	var faculties = assets.GetFaculties(filial)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(faculties)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите факультет").WithOneTimeKeyboard()
}
