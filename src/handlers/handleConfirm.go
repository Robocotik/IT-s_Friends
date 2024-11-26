package handle

import (
	"fmt"
	keyboard "github.com/Robocotik/IT-s_Friends/src/components/keyboards"
	"github.com/Robocotik/IT-s_Friends/src/components/structures"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleConfirm(bot *telego.Bot, msg telego.Message, identity *structures.Identity, isMe bool) {
	var phrase string
	if isMe {
		phrase = "Выходит ты учишься в %s в группе  %s, верно?"
			
	} else {
		phrase = "Выходит твой друг учится в %s в группе  %s, верно?"
			
	}
	keyboard := keyboard.CreateKeyboardConfirm()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf(phrase, identity.Filial, identity.Group),
	).WithReplyMarkup(keyboard))
}
