package logic

import (
	"Friends/src/assets"
	"Friends/src/utils"
	"Friends/src/components/structures"
	"Friends/src/handlers"

	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	StateStart structures.State = iota
	StateDefault
	StateAskFilial
	StateAskCourse
	StateAskFaculty
	StateAskCathedra
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
		user.Filial = utils.ParseString(bot, msg, "филиал", filials[:])
		handle.HandleSelectCourse(bot, msg, user.Filial)
		user.State = StateAskCourse

	case StateAskCourse:
		var courses = assets.GetCourses(user.Filial)
		user.Course = utils.ParseString(bot, msg, "курс", courses)
		handle.HandleSelectFaculty(bot, msg, user.Filial)
		user.State = StateAskFaculty

	case StateAskFaculty:
		var faculties = assets.GetFaculties(user.Filial)
		user.Faculty = utils.ParseString(bot, msg, "факультет", faculties)
		handle.HandleSelectCathedra(bot, msg, user.Filial,  user.Faculty)
		user.State = StateAskCathedra

	case StateAskCathedra:
		var cathedras = assets.GetCathedras(user.Filial, user.Faculty)
		user.Cathedra = utils.ParseString(bot, msg, "кафедра", cathedras)
		user.State = StateAskGroup
		handle.HandleSelectGroup(bot, msg, user.Filial, user.Faculty, user.Course, user.Cathedra)

	case StateAskGroup:

		user.Group = utils.ParseString(bot, msg, "группа", assets.Group[:])
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

		SearchGroupUID(bot, msg, user)

	default:
		_, _ = bot.SendMessage(tu.Message(
			msg.Chat.ChatID(),
			fmt.Sprintf("Неизвестная команда"),
		))
		panic("unknown state")
	}
}
