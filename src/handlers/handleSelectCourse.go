package handle

import (
	"Friends/src/components/keyboards"
	// "Friends/src/messages"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"fmt"
)


func HandleSelectCourse(bot *telego.Bot, msg telego.Message) {
	bot.EditMessageReplyMarkup(nil);
	keyboard := keyboard.CreateKeyboardCourse()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите курс"),
	).WithReplyMarkup(keyboard))
}