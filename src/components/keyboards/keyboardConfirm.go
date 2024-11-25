package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/src/components/structures"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardConfirm() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(structures.YES),
			tu.KeyboardButton(structures.NO),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder(fmt.Sprintf("%s / %s", structures.YES, structures.NO)).WithOneTimeKeyboard()
}
