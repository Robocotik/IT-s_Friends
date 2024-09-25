package keyboard

import (
	"Friends/src/messages"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardCourse() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow( // Row 1
			tu.KeyboardButton(messages.Years[0]),
			tu.KeyboardButton(messages.Years[1]),
		),
		tu.KeyboardRow( // Row 2
			tu.KeyboardButton(messages.Years[2]),
			tu.KeyboardButton(messages.Years[3]),
		),
		tu.KeyboardRow( // Row 2
			tu.KeyboardButton(messages.Years[4]),
			tu.KeyboardButton(messages.Years[5]),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите курс")
}
