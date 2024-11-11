package keyboard

import (
	"Friends/src/assets"
	"Friends/src/utils"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardCourse(filial string, faculty string, cathedra string) *telego.ReplyKeyboardMarkup {
	var courses = assets.GetCourses(filial, faculty, cathedra)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(courses)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите курс").WithOneTimeKeyboard()
}
