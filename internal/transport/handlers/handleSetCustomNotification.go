package handle

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleSetCustomNotification(bot *telego.Bot, msg telego.Message) {
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Укажите, когда вам удобно получать рассылку.\nФормат часы:минуты (например 00:05)"),
	))
}
