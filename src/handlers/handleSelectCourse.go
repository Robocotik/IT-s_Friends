package handle

import (
	keyboard "Friends/src/components/keyboards"
	"Friends/src/components/structures"

	// "Friends/src/messages"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)


func HandleSelectCourse(conn *pgx.Conn, bot *telego.Bot, msg telego.Message, identity *structures.Identity) {
	keyboard := keyboard.CreateKeyboardCourse(conn, bot, msg, identity)
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Выберите курс"),
	).WithReplyMarkup(keyboard))
	
}