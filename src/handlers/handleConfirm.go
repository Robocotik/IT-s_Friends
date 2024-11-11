package handle

import (
	"Friends/src/components/keyboards"
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
			"Выходит твой друг учится в %sом филиале на %s, на %s%s, верно?",
			filial[:len(filial)-2], user.Course, user.Faculty, user.Cathedra),
	).WithReplyMarkup(keyboard))
}
