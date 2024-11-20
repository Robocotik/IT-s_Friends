package utils

import (
	// "Friends/src/components/structures"
	"Friends/src/components/structures"
	"Friends/src/entities"
	"strconv"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowLesson(bot *telego.Bot, msg telego.Message, lesson entities.IDay, showDayName bool, isCh bool, isLast bool) {
	dataToShowBold := "-1"
	dataToShow := "-1"
	dataToShowCabinet := ""
	if len(lesson.Audiences) > 0 {
		for _, audience := range lesson.Audiences {
			dataToShowCabinet += audience.Name + ", "
		}
	}
	if dataToShowCabinet == ""{
		dataToShowCabinet = "Кабинет не указан"
	}
	if (isCh && lesson.Week != "zn") || (!isCh && lesson.Week != "ch") {
		dataToShow = "📅 " + strconv.Itoa(lesson.Time) + " пара ( " + lesson.StartTime + " - " + lesson.EndTime + ")\n\n"
		dataToShowBold = "🎓 " + (lesson.Discipline.FullName) + "\n\n"
		dataToShowCabinet = "🚪 " + dataToShowCabinet
	}
	if dataToShowBold == "-1" {
		return
	}
	dayPhrase := ""
	if showDayName {
		dayPhrase = GetPhrase(lesson.Day) + "\n\n\n"
	}
	_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(msg.Chat.ID),
		tu.Entity(dayPhrase).Underline(),
		tu.Entity(dataToShow),
		tu.Entity(dataToShowBold).Bold(),
		tu.Entity(dataToShowCabinet),
	))
	if !isLast {
		_, _ = bot.SendMessage(tu.MessageWithEntities(tu.ID(msg.Chat.ID),
			tu.Entity(structures.BorderMinus),
		))
	}

}
