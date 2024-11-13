package keyboard

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardAddToFavourite() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("Вернуться к поиску"),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Вернуться к поиску").WithOneTimeKeyboard()
}
