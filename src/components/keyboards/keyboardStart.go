package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/src/messages"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardStart() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(messages.Start),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("/start").WithOneTimeKeyboard()
}
