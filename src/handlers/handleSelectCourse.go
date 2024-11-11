package handle

import (
	keyboard "Friends/src/components/keyboards"
	// "Friends/src/messages"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)


func HandleSelectCourse(bot *telego.Bot, msg telego.Message, filial string) {
	keyboard := keyboard.CreateKeyboardCourse(filial)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите курс"),
	).WithReplyMarkup(keyboard))
	
}