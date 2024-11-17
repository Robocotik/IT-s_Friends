package logic

import (
	keyboard "Friends/src/components/keyboards"
	"Friends/src/entities"
	"Friends/src/components/structures"
	"Friends/src/utils"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowTimetable(bot *telego.Bot, msg telego.Message, request entities.Final_timetable, ch_zn_selected string) {
	prevDay := request.Data.Schedule[0].Day
	for index, lesson := range request.Data.Schedule {
		utils.ShowLesson(bot, msg, lesson, index == 0 || prevDay != lesson.Day, ch_zn_selected == "числитель", index == len(request.Data.Schedule) - 1 )
		prevDay = lesson.Day
	}

	keyboard := keyboard.CreateKeyboardReturnToSearch() // что - то придумать
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf(structures.BorderMinus),
	).WithReplyMarkup(keyboard))

}
