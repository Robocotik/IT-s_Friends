package handle

import (
	"Friends/src/components/keyboards"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"fmt"
)


func HandleSelectFaculty(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardFaculty()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите факультет"),
	).WithReplyMarkup(keyboard))
}