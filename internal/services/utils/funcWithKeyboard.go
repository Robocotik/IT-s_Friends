package utils

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func FuncWithKeyboard(bot *telego.Bot, msg telego.Message, fn func() (string, error), keyboard *telego.ReplyKeyboardMarkup) {
	val, _ := fn()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		val,
	).WithReplyMarkup(keyboard.WithOneTimeKeyboard()))

}
