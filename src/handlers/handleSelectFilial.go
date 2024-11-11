package handle

import (
	keyboard "Friends/src/components/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)


func HandleSelectFilial(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardFilial()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите филиал"),
	).WithReplyMarkup(keyboard))
}