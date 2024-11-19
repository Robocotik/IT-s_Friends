package utils

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func RiseError(bot *telego.Bot, msg telego.Message, err error) {
	if err != nil {
		_, _ = bot.SendMessage(tu.Message(
			msg.Chat.ChatID(),
			fmt.Sprintf("Неверный %s, попробуй еще раз :)", err.Error()),
		))
	}
}
