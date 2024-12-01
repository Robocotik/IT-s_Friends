package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/assets/consts"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardGroupFound() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(consts.SHOW_SCHEDULE),
			tu.KeyboardButton(consts.ADD_TO_FAVOURITE),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder(consts.SHOW_SCHEDULE + " / " + consts.ADD_TO_FAVOURITE).WithOneTimeKeyboard()
}
