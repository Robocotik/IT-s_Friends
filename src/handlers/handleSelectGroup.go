package handle

import (
	keyboard "Friends/src/components/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)


func HandleSelectGroup(bot *telego.Bot, msg telego.Message, filial string, faculty string, course string,  cathedra string ) {
	keyboard := keyboard.CreateKeyboardGroup(filial, course, faculty, cathedra)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите группу"),
	).WithReplyMarkup(keyboard))
}