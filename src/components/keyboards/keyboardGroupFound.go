package keyboard

import (
	"Friends/src/components/structures"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardGroupFound() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(structures.SHOW_SCHEDULE),
			tu.KeyboardButton(structures.ADD_TO_FAVOURITE),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder(structures.SHOW_SCHEDULE + " / " + structures.ADD_TO_FAVOURITE).WithOneTimeKeyboard()
}
