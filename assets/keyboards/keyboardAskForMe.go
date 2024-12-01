package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/assets/consts"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardAskForMe() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(consts.YES),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder(consts.YES).WithOneTimeKeyboard()
}
