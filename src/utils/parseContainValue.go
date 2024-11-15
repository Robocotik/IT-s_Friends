package utils

import (
	"strings"

	"github.com/mymmrac/telego"
)

func ParseContainString(bot *telego.Bot, msg telego.Message, err string, possibleData []string) string {
	found := false
	if msg.Text == "" {
		return "Я встретил пустую строку"
	}
	for _, data := range possibleData {
		if strings.Contains(msg.Text, data) {
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
