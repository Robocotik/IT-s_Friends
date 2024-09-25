package handleFaculty

import (
	"Friends/src/components/keyboards"
	// "Friends/src/messages"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"fmt"
)


func HandleSelectFaculty(bot *telego.Bot, update telego.Update) {
	bot.EditMessageReplyMarkup(nil);
	keyboard := keyboard.CreateKeyboardFaculty()
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Выберите факультет"),
	).WithReplyMarkup(keyboard))
}