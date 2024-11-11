package handle

import (
	keyboard "Friends/src/components/keyboards"
	// "Friends/src/messages"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)


func HandleSelectCourse(bot *telego.Bot, msg telego.Message, filial string, faculty string, cathedra string) {
	keyboard := keyboard.CreateKeyboardCourse(filial, faculty, cathedra)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите курс"),
	).WithReplyMarkup(keyboard))
	
}