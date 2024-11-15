package utils

import (
	"Friends/src/components/structures"
	"Friends/src/entities"
	"strconv"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowDay(bot *telego.Bot, msg telego.Message, day entities.IDay, showDayName bool, isCh bool) {

	if showDayName {
		_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(msg.Chat.ID),
			tu.Entity(structures.BorderMinus + GetPhrase(day.Day)).Underline()),
		)
	}

	_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(msg.Chat.ID),
		tu.Entity(structures.BorderMinus + strconv.Itoa(day.Time) + ")" + day.Discipline.FullName+" : "+day.StartTime+" - "+day.EndTime).Bold()),
	)
}
