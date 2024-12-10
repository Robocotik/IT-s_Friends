package keyboard

import (
	"github.com/Robocotik/IT-s_Friends/internal/database"
	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/utils"

	// "Friends/src/messages"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	// "fmt"
)


func CreateKeyboardCathedra(bd database.IBd, bot *telego.Bot, chatID int64, identity *structures.Identity) *telego.ReplyKeyboardMarkup {
	var cathedras = bd.GetCathedras(bot, chatID, identity)
	var items_rows [][]telego.KeyboardButton = utils.GetItemsRow(cathedras)
	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите кафедру").WithOneTimeKeyboard()
}
