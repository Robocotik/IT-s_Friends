package keyboard

import (
	// "Friends/src/assets/faculty"
	"Friends/src/messages"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	// "fmt"
)

func CreateKeyboardFaculty() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow( // Row 1
			tu.KeyboardButton(messages.Years[4]),
			tu.KeyboardButton(messages.Years[3]),
		),
		tu.KeyboardRow( // Row 2
			tu.KeyboardButton(messages.Years[1]),
			tu.KeyboardButton(messages.Years[2]),
		),
		tu.KeyboardRow( // Row 2
			tu.KeyboardButton(messages.Years[5]),
			tu.KeyboardButton(messages.Years[2]),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите курс")
}
