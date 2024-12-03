package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/assets/consts"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardStartMenu() *telego.ReplyKeyboardMarkup {
	tu.InlineKeyboard()
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(consts.FIND_NEW_FRIENDS),
			tu.KeyboardButton(consts.SHOW_FRIENDS),
		),
		tu.KeyboardRow(
			tu.KeyboardButton(consts.SET_NOTIFICATIONS),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder(consts.FIND_NEW_FRIENDS + " / " + consts.SHOW_FRIENDS + " / " + consts.SET_NOTIFICATIONS).WithOneTimeKeyboard()
}
