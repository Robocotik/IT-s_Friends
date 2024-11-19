package utils

import (
	"github.com/mymmrac/telego"
)

func ParseString(bot *telego.Bot, msg telego.Message, err error, possibleData []string) string {
	found := false
	if msg.Text == "" {
		return "Я встретил пустую строку"
	}
	for _, data := range possibleData {
		if msg.Text == data {
			found = true
			break
		}
	}
	if !found {
		RiseError(bot, msg, err)
		return "-1"
	}
	return msg.Text
}
