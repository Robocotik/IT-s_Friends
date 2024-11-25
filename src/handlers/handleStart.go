package handle

import (
	keyboard "github.com/Robocotik/IT-s_Friends/src/components/keyboards"
	"github.com/Robocotik/IT-s_Friends/src/messages"
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
