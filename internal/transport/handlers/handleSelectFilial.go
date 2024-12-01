package handle

import (
	keyboard "github.com/Robocotik/IT-s_Friends/assets/keyboards"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleSelectFilial(conn *pgx.Conn, bot *telego.Bot, msg telego.Message) {
	keyboard := keyboard.CreateKeyboardFilial(conn, bot, msg)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите филиал"),
	).WithReplyMarkup(keyboard))
}
