package logic

import (
	"Friends/src/entities"
	"Friends/src/utils"

	"github.com/mymmrac/telego"
)

func ShowTimetable(bot *telego.Bot, msg telego.Message, request entities.Final_timetable) {
	prevDay := request.Data.Schedule[0].Day
	for index, day := range request.Data.Schedule {
		utils.ShowDay(bot, msg, day, index == 0 || prevDay != day.Day)
		prevDay = day.Day
	}
}
