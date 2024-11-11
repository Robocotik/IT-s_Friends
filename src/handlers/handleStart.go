package handle

import (
	keyboard "Friends/src/components/keyboards"
	"Friends/src/messages"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleStart(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardStart()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf(messages.HelloPhrase),
	).WithReplyMarkup(keyboard))
}
