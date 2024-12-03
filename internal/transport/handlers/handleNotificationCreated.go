package handle

import (
	"fmt"
	keyboard "github.com/Robocotik/IT-s_Friends/assets/keyboards"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleNotificationCreated(bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardReturnToSearch()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Уведомления успешно добавлены.\nТеперь в выбранное время вы будете получать информацию о том, с кем из друзей можете поехать в ВУЗ вместе."),
	).WithReplyMarkup(keyboard))
}
