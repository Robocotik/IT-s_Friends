package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/internal/database"
	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/utils"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardCourse(bd database.IBd, bot *telego.Bot, chatID int64, identity *structures.Identity) *telego.ReplyKeyboardMarkup {
	var courses = bd.GetCourses(bot, chatID, identity)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(courses)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите курс").WithOneTimeKeyboard()
}
