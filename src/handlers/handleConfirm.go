package handle

import (
	keyboard "Friends/src/components/keyboards"
	"Friends/src/components/structures"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleConfirm(bot *telego.Bot, msg telego.Message, user *structures.User) {
	keyboard := keyboard.CreateKeyboardConfirm()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),

		fmt.Sprintf(
			"Выходит твой друг учится в %s на %s курсе, на %s, верно?",
			user.Filial, user.Course, user.Group),
	).WithReplyMarkup(keyboard))
}
