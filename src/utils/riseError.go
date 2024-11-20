package utils

import (
	"github.com/mymmrac/telego"
)

func RiseError(bot *telego.Bot, msg telego.Message, err error) {
	if err != nil {
		WriteMessage(bot, msg, "Неверный "+err.Error()+", попробуй еще раз :)")
	}
}
