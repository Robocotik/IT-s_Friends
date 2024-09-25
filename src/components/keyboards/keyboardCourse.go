package keyboard

import (
	"Friends/src/assets/emoji"
	"Friends/src/assets"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardCourse() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow( // Row 1
			tu.KeyboardButton(emoji.Courses[0]+assets.Courses[0]),
			tu.KeyboardButton(emoji.Courses[1]+assets.Courses[1]),
		),
		tu.KeyboardRow( // Row 2
			tu.KeyboardButton(emoji.Courses[2]+assets.Courses[2]),
			tu.KeyboardButton(emoji.Courses[3]+assets.Courses[3]),
		),
		tu.KeyboardRow( // Row 2
			tu.KeyboardButton(emoji.Courses[4]+assets.Courses[4]),
			tu.KeyboardButton(emoji.Courses[5]+assets.Courses[5]),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите курс")
}
