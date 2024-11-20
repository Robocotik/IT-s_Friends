package handle

import (
	keyboard "Friends/src/components/keyboards"
	"Friends/src/components/structures"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleSelectFaculty(bot *telego.Bot, msg telego.Message, friend *structures.AskedFriend ) {
	keyboard := keyboard.CreateKeyboardFaculty(friend.Filial)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите факультет"),
	).WithReplyMarkup(keyboard))
}
