package utils

import (
	"Friends/src/entities"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoutil"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowDay(bot *telego.Bot, msg telego.Message, day []entities.IDay, isCh bool, keyboard *telego.ReplyKeyboardMarkup) {
	var entities []telegoutil.MessageEntityCollection
	if len(day) > 0 {
		entities = append(entities, tu.Entity(GetPhrase(day[0].Day)+"\n\n").Underline())
	}

	for _, lesson := range day {
		entities = append(entities, ShowLesson(msg, lesson, isCh)...)
	}
	_, _ = bot.SendMessage(tu.MessageWithEntities(msg.Chat.ChatID(), entities...).WithReplyMarkup(keyboard))

}
