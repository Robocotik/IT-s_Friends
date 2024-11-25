package keyboard

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardAddMe() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("Да"),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Да").WithOneTimeKeyboard()
}
