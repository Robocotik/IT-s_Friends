package handleStart

import (
	"Friends/src/components/keyboard"
	"Friends/src/messages"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"fmt"
)


func HandleStartCommand(bot *telego.Bot, update telego.Update) {
	keyboard := keyboard.CreateKeyboard()
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(messages.HelloPhrase),
	).WithReplyMarkup(keyboard))
}
