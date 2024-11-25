package handle

import (
	keyboard "github.com/Robocotik/IT-s_Friends/src/components/keyboards"
	"github.com/Robocotik/IT-s_Friends/src/components/structures"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleConfirm(bot *telego.Bot, msg telego.Message, identity *structures.Identity) {
	keyboard := keyboard.CreateKeyboardConfirm()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),

		fmt.Sprintf(
			"Выходит твой друг учится в %s на %s курсе, на %s, верно?",
			identity.Filial, identity.Course, identity.Group),
	).WithReplyMarkup(keyboard))
}
