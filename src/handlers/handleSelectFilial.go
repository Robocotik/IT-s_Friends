package handle

import (
	"Friends/src/components/keyboards"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"fmt"
)


func HandleSelectFilial(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardFilial()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите филиал"),
	).WithReplyMarkup(keyboard))
}