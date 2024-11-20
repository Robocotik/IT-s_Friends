package utils

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func WriteMessage(bot *telego.Bot, msg telego.Message, text string) {
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		text,
	))
}
