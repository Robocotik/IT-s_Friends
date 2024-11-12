package logic

import (
	"Friends/src/components/structures"
	"Friends/src/utils"

	"github.com/mymmrac/telego"
)

func ShowTimetable(bot *telego.Bot, msg telego.Message, request structures.Final_timetable) {
	for _, day := range request.Data.Schedule {
		utils.ShowDay(bot, msg, day)
	}
}
