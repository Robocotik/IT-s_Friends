package handle

import (
	keyboard "Friends/src/components/keyboards"
	"Friends/src/components/structures"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleSelectGroup(bot *telego.Bot, msg telego.Message, friend *structures.AskedFriend) {
	keyboard := keyboard.CreateKeyboardGroup(friend.Filial, friend.Course, friend.Faculty, friend.Cathedra)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите группу"),
	).WithReplyMarkup(keyboard))
}
