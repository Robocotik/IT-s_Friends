package output

import (
	// "Friends/src/components/structures"

	"github.com/Robocotik/IT-s_Friends/internal/models/entities"
	"strconv"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowLesson(msg telego.Message, lesson entities.IDay, isCh bool) []tu.MessageEntityCollection {
	dataToShowBold := "-1"
	dataToShow := "-1"
	dataToShowCabinet := ""
	if len(lesson.Audiences) > 0 {
		for _, audience := range lesson.Audiences {
			dataToShowCabinet += audience.Name + ", "
		}
	}
	if dataToShowCabinet == "" {
		dataToShowCabinet = "Кабинет не указан"
	}
	if (isCh && lesson.Week != "zn") || (!isCh && lesson.Week != "ch") {
		dataToShow = "📅 " + strconv.Itoa(lesson.Time) + " пара ( " + lesson.StartTime + " - " + lesson.EndTime + " )\n"
		dataToShowBold = "🎓 " + (lesson.Discipline.FullName) + "\n"
		dataToShowCabinet = "🚪 " + dataToShowCabinet[:len(dataToShowCabinet)-2] + "\n\n"
	}
	if dataToShowBold == "-1" {
		return nil
	}

	res := []tu.MessageEntityCollection{
		
		tu.Entity(dataToShow),
		tu.Entity(dataToShowBold).Bold(),
		tu.Entity(dataToShowCabinet),
	}
	return res

}
