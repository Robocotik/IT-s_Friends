package handle

import (
	keyboard "Friends/src/components/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)


func HandleSelectCathedra(bot *telego.Bot, msg telego.Message, filial string, faculty string) {
	keyboard := keyboard.CreateKeyboardCathedra(filial, faculty)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите кафедру"),
	).WithReplyMarkup(keyboard))
}