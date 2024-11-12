package handle

import (
	keyboard "Friends/src/components/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleGroupNotFound(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardGroupNotFound()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("🤨 Видимо расписание для группы еще не составлено..."),
	).WithReplyMarkup(keyboard))
}
