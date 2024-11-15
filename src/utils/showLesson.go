package utils

import (
	// "Friends/src/components/structures"
	"Friends/src/components/structures"
	"Friends/src/entities"
	"strconv"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowLesson(bot *telego.Bot, msg telego.Message, lesson entities.IDay, showDayName bool, isCh bool) {
	dataToShowBold := "-1"
	dataToShow := "-1"
	if (isCh && lesson.Week != "zn") || (!isCh && lesson.Week != "ch") {
		dataToShow = "📅 " + strconv.Itoa(lesson.Time) + " пара ( " + lesson.StartTime + " - " + lesson.EndTime + ") :\n"
		dataToShowBold = "🎓 " + (lesson.Discipline.FullName)

	}
	if dataToShowBold == "-1" {
		return
	}

	if showDayName {
		_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(msg.Chat.ID),
			tu.Entity(GetPhrase(lesson.Day)).Underline()),
		)
	}
	_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(msg.Chat.ID),
		tu.Entity(structures.BorderMinus),
	))
	_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(msg.Chat.ID),
		tu.Entity(dataToShow),
		tu.Entity(dataToShowBold).Bold(),
	))
}
