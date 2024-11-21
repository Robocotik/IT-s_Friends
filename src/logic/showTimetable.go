package logic

import (
	keyboard "Friends/src/components/keyboards"
	"Friends/src/entities"
	"Friends/src/utils"
	"github.com/mymmrac/telego"
)

func ShowTimetable(bot *telego.Bot, msg telego.Message, request entities.Final_timetable, ch_zn_selected string) {
	var lessons_on_day []entities.IDay
	keyboard := keyboard.CreateKeyboardReturnToSearch()
	prevDay := request.Data.Schedule[0].Day
	for _, lesson := range request.Data.Schedule {
		if lesson.Day != prevDay {
			utils.ShowDay(bot, msg, lessons_on_day, ch_zn_selected == "числитель", keyboard)
			lessons_on_day = []entities.IDay{}
			lessons_on_day = append(lessons_on_day, lesson)
			prevDay = lesson.Day
			continue
		}
		lessons_on_day = append(lessons_on_day, lesson)
	}
	utils.ShowDay(bot, msg, lessons_on_day, ch_zn_selected == "числитель", keyboard)

}
