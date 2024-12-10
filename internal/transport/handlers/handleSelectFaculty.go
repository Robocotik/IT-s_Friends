package handle

import (
	"fmt"

	keyboard "github.com/Robocotik/IT-s_Friends/assets/keyboards"
	"github.com/Robocotik/IT-s_Friends/internal/database"
	"github.com/Robocotik/IT-s_Friends/internal/models/structures"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleSelectFaculty(bd database.IBd, bot *telego.Bot, msg telego.Message, identity *structures.Identity) {
	keyboard := keyboard.CreateKeyboardFaculty(bd, bot, msg.Chat.ID, identity)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите факультет"),
	).WithReplyMarkup(keyboard))
}
