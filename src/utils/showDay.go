package utils

import (
	"Friends/src/components/structures"
	"strconv"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowDay(bot *telego.Bot, msg telego.Message, day structures.IDay, showDayName bool) {

	if showDayName {
		_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(msg.Chat.ID),
			tu.Entity(GetPhrase(day.Day)).Underline()),
		)
	}

	_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(msg.Chat.ID),
		tu.Entity(strconv.Itoa(day.Time) + ")" + day.Discipline.FullName+" : "+day.StartTime+" - "+day.EndTime).Bold()),
	)
}
