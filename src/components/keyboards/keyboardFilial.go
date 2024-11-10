package keyboard

import (
	"Friends/src/assets"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardFilial() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(assets.Filials[0]),
		),
		tu.KeyboardRow(
			tu.KeyboardButton(assets.Filials[1]),
		),
		tu.KeyboardRow(
			tu.KeyboardButton(assets.Filials[2]),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите филиал").WithOneTimeKeyboard()
}
