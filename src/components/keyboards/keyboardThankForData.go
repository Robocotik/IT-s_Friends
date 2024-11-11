package keyboard

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardThankForData() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("Начать поиск 🔎"),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Начать поиск 🔎").WithOneTimeKeyboard()
}