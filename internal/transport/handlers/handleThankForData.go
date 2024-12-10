package handle

import (
	keyboard "github.com/Robocotik/IT-s_Friends/assets/keyboards"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleThankForData(bot *telego.Bot, chatID int64) {
	keyboard := keyboard.CreateKeyboardThankForData()
	_, _ = bot.SendMessage(tu.Message(
		telego.ChatID{ID: chatID},
		fmt.Sprintf("Спасибо за информацию!"),
	).WithReplyMarkup(keyboard))
}