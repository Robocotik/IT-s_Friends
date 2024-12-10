package output

import (
	"github.com/mymmrac/telego"
)

func RiseError(bot *telego.Bot, chatID int64, err error) {
	if err != nil {
		WriteMessage(bot, chatID, "Неверный "+err.Error()+", попробуй еще раз :)")
	}
}
