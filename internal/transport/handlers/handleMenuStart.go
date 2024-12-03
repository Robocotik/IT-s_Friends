package handle

import (
	keyboard "github.com/Robocotik/IT-s_Friends/assets/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleMenuStart(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardStartMenu()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите действие"),
	).WithReplyMarkup(keyboard))
}
