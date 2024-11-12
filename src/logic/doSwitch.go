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

const (
	StateStart structures.State = iota
	StateDefault
	StateAskFilial
	StateAskFaculty
	StateAskCathedra
	StateAskCourse
	StateAskGroup
	StateConfirm
	StateSearch
)

func DoSwitch(user *structures.User, bot *telego.Bot, msg telego.Message) {
	switch user.State {

	case StateStart:
		msg.Text = ""
		handle.HandleStart(bot, msg)
		user.State = StateDefault
		msg.Text = ""

	case StateDefault:
		user.Filial = utils.ParseString(bot, msg, "Погнали", []string{"Погнали"})
		// if user.Filial == "-1" {
		// 	user.State = StateStart
		// 	break
		// }
		handle.HandleSelectFilial(bot, msg)
		user.State = StateAskFilial

	case StateAskFilial:
		var filials = assets.GetFilials()
		user.Filial = utils.ParseString(bot, msg, "филиал", filials)
		handle.HandleSelectFaculty(bot, msg, user.Filial)
		user.State = StateAskFaculty

	case StateAskFaculty:
		var faculties = assets.GetFaculties(user.Filial)
		user.Faculty = utils.ParseString(bot, msg, "факультет", faculties)
		handle.HandleSelectCathedra(bot, msg, user.Filial, user.Faculty)
		user.State = StateAskCathedra

	case StateAskCathedra:
		var cathedras = assets.GetCathedras(user.Filial, user.Faculty)
		user.Cathedra = utils.ParseString(bot, msg, "кафедра", cathedras)
		user.State = StateAskCourse
		handle.HandleSelectCourse(bot, msg, user.Filial, user.Faculty, user.Cathedra)

	case StateAskCourse:
		var courses = assets.GetCourses(user.Filial, user.Faculty, user.Cathedra)
		user.Course = utils.ParseString(bot, msg, "курс", courses)
		handle.HandleSelectGroup(bot, msg, user.Filial, user.Faculty, user.Course, user.Cathedra)
		user.State = StateAskGroup

	case StateAskGroup:
		var groups = assets.GetGroups(user.Filial, user.Course, user.Faculty, user.Cathedra)
		user.Group = utils.ParseString(bot, msg, "-", groups)
		user.State = StateConfirm
		handle.HandleConfirm(bot, msg, user)

	case StateConfirm:

		if msg.Text == structures.YES {
			handle.HandleThankForData(bot, msg)
			user.State = StateSearch

		} else {
			handle.HandleSelectFilial(bot, msg)
			user.State = StateAskFilial
		}

	case StateSearch:

		uid := SearchGroupUID(bot, msg, user)
		request := DoRequest(uid)
		ShowTimetable(bot, msg, request)

	default:
		_, _ = bot.SendMessage(tu.Message(
			msg.Chat.ChatID(),
			fmt.Sprintf("Неизвестная команда"),
		))
		panic("unknown state")
	}
}
