package keyboard

import (
	"Friends/src/assets"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardCourse() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(assets.Courses[0]),
			tu.KeyboardButton(assets.Courses[1]),
		),
		tu.KeyboardRow(
			tu.KeyboardButton(assets.Courses[2]),
			tu.KeyboardButton(assets.Courses[3]),
		),
		tu.KeyboardRow(
			tu.KeyboardButton(assets.Courses[4]),
			tu.KeyboardButton(assets.Courses[5]),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите курс").WithOneTimeKeyboard()
}
