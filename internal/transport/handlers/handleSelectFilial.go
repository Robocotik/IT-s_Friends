package handle

import (
	"fmt"

	keyboard "github.com/Robocotik/IT-s_Friends/assets/keyboards"
	"github.com/Robocotik/IT-s_Friends/internal/database"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleSelectFilial(bd database.IBd, bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardFilial(bd, bot, msg)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите филиал"),
	).WithReplyMarkup(keyboard))
}
