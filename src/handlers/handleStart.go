package handle

import (
	"Friends/src/components/keyboards"
	"Friends/src/messages"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleStartCommand(bot *telego.Bot, update telego.Update) {
	bot.EditMessageReplyMarkup(nil);
	keyboard := keyboard.CreateKeyboardStart()
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(messages.HelloPhrase),
	).WithReplyMarkup(keyboard))
}
