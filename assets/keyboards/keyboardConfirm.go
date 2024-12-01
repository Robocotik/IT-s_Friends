package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/assets/consts"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardConfirm() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(consts.YES),
			tu.KeyboardButton(consts.NO),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder(fmt.Sprintf("%s / %s", consts.YES, consts.NO)).WithOneTimeKeyboard()
}
