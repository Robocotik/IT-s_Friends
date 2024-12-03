package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/assets/consts"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardSetNotifications() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(consts.H1_BEFORE),
		),
		tu.KeyboardRow(
			tu.KeyboardButton(consts.H2_BEFORE),
		),
		tu.KeyboardRow(
			tu.KeyboardButton(consts.H3_BEFORE),
		),
		tu.KeyboardRow(
			tu.KeyboardButton(consts.CUSTOM_TIME),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Укажите, когда вам удобно получать рассылку").WithOneTimeKeyboard()
}
