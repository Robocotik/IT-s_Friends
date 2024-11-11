package handle

import (
	keyboard "Friends/src/components/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleThankForData(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardThankForData()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Спасибо за информацию!"),
	).WithReplyMarkup(keyboard))
}