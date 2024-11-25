package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/src/components/structures"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardStartMenu() *telego.ReplyKeyboardMarkup {
	tu.InlineKeyboard()
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(structures.FIND_NEW_FRIENDS),
			tu.KeyboardButton(structures.SHOW_FRIENDS),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder(structures.FIND_NEW_FRIENDS + " / " + structures.SHOW_FRIENDS).WithOneTimeKeyboard()
}
