package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/src/utils"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardShowTimetable() *telego.ReplyKeyboardMarkup {
	phrases := utils.GetChZnPhrases(utils.GetChZn())
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(phrases[0]),
			tu.KeyboardButton(phrases[1]),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder(phrases[0] + " / " + phrases[1]).WithOneTimeKeyboard()
}
