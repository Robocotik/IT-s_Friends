package keyboard

import (
	"Friends/src/assets"
	"Friends/src/utils"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardCourse(faculty string) *telego.ReplyKeyboardMarkup {
	var courses = assets.GetCourses(faculty)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(courses)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите курс").WithOneTimeKeyboard()
}
