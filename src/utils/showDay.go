package utils

import (
	"Friends/src/components/structures"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowDay(bot *telego.Bot, msg telego.Message, day structures.IDay) {
	fmt.Println("МНЕ ПРИСЛАЛИ ДЕНЬ", day)
	_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(msg.Chat.ID),
		tu.Entity(GetPhrase(day.Day)).Underline()),
	)
	_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(msg.Chat.ID),
		tu.Entity("бла бла бла").Bold()),
	)
}
