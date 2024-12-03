package input

import (
	"errors"

	"github.com/mymmrac/telego"
)

func ParseStringOrPeriod(bot *telego.Bot, msg telego.Message, possibleData []string) (string, error) { // return val and flag isCustom
	value, err := ParseContainString(bot, msg, possibleData)
	if err != nil {
		return value, nil
	}
	value, err = CheckPeriod(msg.Text)
	if err != nil {
		return value,  nil
	}
	return "", errors.New("incorrect input")
}
