package utils

import (
	"errors"

	"github.com/mymmrac/telego"
)

func ParseString(bot *telego.Bot, msg telego.Message, err error, possibleData []string) (string, error) {
	found := false
	for _, data := range possibleData {
		if msg.Text == data {
			found = true
			break
		}
	}
	if !found {
		RiseError(bot, msg, err)
		return "", errors.New("unknown message")
	}
	return msg.Text, nil
}
