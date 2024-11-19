package utils

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func FuncWithKeyboard(bot *telego.Bot, msg telego.Message, fn func() []string, keyboard *telego.ReplyKeyboardMarkup) {
	for _, res := range fn() {
		_, _ = bot.SendMessage(tu.Message(
			msg.Chat.ChatID(),
			res,
		).WithReplyMarkup(keyboard.WithOneTimeKeyboard()))
	}
}
