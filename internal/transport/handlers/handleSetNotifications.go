package handle

import (
	"fmt"

	keyboard "github.com/Robocotik/IT-s_Friends/assets/keyboards"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleSetNotifications(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardSetNotifications()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Укажите, когда вам удобно получать рассылку."),
	).WithReplyMarkup(keyboard))
}
