package handle

import (
	keyboard "Friends/src/components/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleSelectFaculty(bot *telego.Bot, msg telego.Message, filial string) {
	keyboard := keyboard.CreateKeyboardFaculty(filial)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите факультет"),
	).WithReplyMarkup(keyboard))
}
