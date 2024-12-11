package output

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)


func WriteMessage(bot *telego.Bot, chatId int64, text string) {
	_, _ = bot.SendMessage(tu.Message(
		telego.ChatID{ID: chatId},
		text,
	))
	bot.GetMe()
}
