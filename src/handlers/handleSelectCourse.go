package handle

import (
	keyboard "Friends/src/components/keyboards"
	"Friends/src/components/structures"
	// "Friends/src/messages"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)


func HandleSelectCourse(bot *telego.Bot, msg telego.Message, friend *structures.AskedFriend) {
	keyboard := keyboard.CreateKeyboardCourse(friend.Filial, friend.Faculty, friend.Cathedra)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите курс"),
	).WithReplyMarkup(keyboard))
	
}