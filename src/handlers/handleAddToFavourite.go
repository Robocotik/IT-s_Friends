package handle

import (
	keyboard "Friends/src/components/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleAddToHavourite(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardReturnToSearch()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Я добавил твоего bestie в избранное 😉"),
	).WithReplyMarkup(keyboard))
}
