package handle

import (
	keyboard "Friends/src/components/keyboards"
	"Friends/src/components/structures"
	"fmt"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleConfirm(bot *telego.Bot, msg telego.Message, user *structures.User) {
	filial := strings.Split(user.Filial, " ")[0]
	keyboard := keyboard.CreateKeyboardConfirm()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),

		fmt.Sprintf(
			"Выходит твой друг учится в %s на %s, на %s%s-%s%s, верно?",
			filial[:len(filial)-2], user.Course, user.Faculty, user.Cathedra, user.Course, user.Group),
	).WithReplyMarkup(keyboard))
}
