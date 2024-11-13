package logic

import (
	"Friends/src/assets"
	"Friends/src/components/structures"
	"Friends/src/handlers"
	"Friends/src/utils"

	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)


func DoSwitch(user *structures.User, bot *telego.Bot, msg telego.Message) {
	switch user.State {

	case structures.StateStart:
		msg.Text = ""
		handle.HandleStart(bot, msg)
		user.State = structures.StateDefault
		msg.Text = ""

	case structures.StateDefault:
		user.Filial = utils.ParseString(bot, msg, "Погнали", []string{"Погнали"})
		// if user.Filial == "-1" {
		// 	user.State = StateStart
		// 	break
		// }
		handle.HandleSelectFilial(bot, msg)
		user.State = structures.StateAskFilial

	case structures.StateAskFilial:
		var filials = assets.GetFilials()
		user.Filial = utils.ParseString(bot, msg, "филиал", filials)
		handle.HandleSelectFaculty(bot, msg, user.Filial)
		user.State = structures.StateAskFaculty

	case structures.StateAskFaculty:
		var faculties = assets.GetFaculties(user.Filial)
		user.Faculty = utils.ParseString(bot, msg, "факультет", faculties)
		handle.HandleSelectCathedra(bot, msg, user.Filial, user.Faculty)
		user.State = structures.StateAskCathedra

	case structures.StateAskCathedra:
		var cathedras = assets.GetCathedras(user.Filial, user.Faculty)
		user.Cathedra = utils.ParseString(bot, msg, "кафедра", cathedras)
		user.State = structures.StateAskCourse
		handle.HandleSelectCourse(bot, msg, user.Filial, user.Faculty, user.Cathedra)

	case structures.StateAskCourse:
		var courses = assets.GetCourses(user.Filial, user.Faculty, user.Cathedra)
		user.Course = utils.ParseString(bot, msg, "курс", courses)
		handle.HandleSelectGroup(bot, msg, user.Filial, user.Faculty, user.Course, user.Cathedra)
		user.State = structures.StateAskGroup

	case structures.StateAskGroup:
		var groups = assets.GetGroups(user.Filial, user.Course, user.Faculty, user.Cathedra)
		user.Group = utils.ParseString(bot, msg, "-", groups)
		user.State = structures.StateConfirm
		handle.HandleConfirm(bot, msg, user)

	case structures.StateConfirm:

		if msg.Text == structures.YES {
			handle.HandleThankForData(bot, msg)
			user.State = structures.StateSearch

		} else {
			handle.HandleSelectFilial(bot, msg)
			user.State = structures.StateAskFilial
		}

	case structures.StateSearch:

		uid := SearchGroupUID(bot, msg, user)
		request := DoRequest(bot, msg, uid)
		if len(request.Data.Schedule) != 0 {
			ShowTimetable(bot, msg, request)
		} else {
			handle.HandleGroupNotFound(bot, msg)
			user.State = structures.StateGroupNotFound
		}
	case structures.StateGroupNotFound:

		handle.HandleSelectFilial(bot, msg)
		user.State = structures.StateAskFilial

	default:
		_, _ = bot.SendMessage(tu.Message(
			msg.Chat.ChatID(),
			fmt.Sprintf("Неизвестная команда"),
		))
		panic("unknown state")
	}
}
