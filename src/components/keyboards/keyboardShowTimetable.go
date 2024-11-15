package keyboard

import (
	"Friends/src/utils"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardShowTimetable() *telego.ReplyKeyboardMarkup {
	isCh := utils.GetChZn() == "числитель" // переписать на словарь 
	phrase_ch := "Числитель"
	phrase_zn := "Знаменатель"
	if isCh {
		phrase_ch += " (Сегодня)"
	} else {
		phrase_zn += " (Сегодня)"
	}
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(phrase_ch),
			tu.KeyboardButton(phrase_zn),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder(phrase_ch + " / " + phrase_zn).WithOneTimeKeyboard()
}
