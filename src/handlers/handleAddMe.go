package handle

import (
	keyboard "github.com/Robocotik/IT-s_Friends/src/components/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleAddMe(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardAddMe()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Вижу ты тут впервые) Мне надо узнать немного о тебе, ты не против?)"),
	).WithReplyMarkup(keyboard))
}
