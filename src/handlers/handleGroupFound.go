package handle

import (
	keyboard "github.com/Robocotik/IT-s_Friends/src/components/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleGroupFound(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardGroupFound()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Да, я нашел такую группу.\nХочешь посмотреть расписание или же просто добавить друга в избранное?"),
	).WithReplyMarkup(keyboard))
}
