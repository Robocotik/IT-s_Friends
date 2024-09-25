package utils

import (
	"github.com/mymmrac/telego"
)

func ParseString(bot *telego.Bot, msg telego.Message, err string, possibleData []string) string {
	found := false

	for _, data := range possibleData {
		if msg.Text == data { // Предполагаю, что вы хотите сравнить текст сообщения
			found = true
			break
		}
	}
	if !found {
		RiseError(bot, msg, err)
	}
	return msg.Text
}
