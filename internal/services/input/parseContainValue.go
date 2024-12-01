package input

import (
	"strings"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"github.com/mymmrac/telego"
)

func ParseContainString(bot *telego.Bot, msg telego.Message, err error, possibleData []string) string {
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
		output.RiseError(bot, msg, err)
		return "-1"
	}
	return msg.Text
}
