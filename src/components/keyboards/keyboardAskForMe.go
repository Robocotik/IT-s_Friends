package keyboard

import (
	"Friends/src/components/structures"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardAskForMe() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(structures.YES),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder(structures.YES).WithOneTimeKeyboard()
}
