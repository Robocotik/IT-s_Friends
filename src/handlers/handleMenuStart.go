package handle

import (
	keyboard "Friends/src/components/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleMenuStart(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardStartMenu()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Хочешь добавить нового друга или посмотреть список уже добавленных?"),
	).WithReplyMarkup(keyboard))
}
