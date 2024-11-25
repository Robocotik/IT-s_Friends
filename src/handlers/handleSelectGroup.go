package handle

import (
	keyboard "Friends/src/components/keyboards"
	"Friends/src/components/structures"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleSelectGroup(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, friend *structures.AskedFriend) {
	keyboard := keyboard.CreateKeyboardGroup(conn, bot, msg, friend)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите группу"),
	).WithReplyMarkup(keyboard))
}
