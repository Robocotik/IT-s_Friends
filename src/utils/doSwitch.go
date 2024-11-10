package utils

import (
	"Friends/src/assets"
	"Friends/src/components/structures"
	"Friends/src/handlers"
	"fmt"
	"strings"

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
		user.Filial = ParseString(bot, msg, "Погнали", []string{"Погнали"})
		// if user.Filial == "-1" {
		// 	user.State = StateStart
		// 	break
		// }
		handle.HandleSelectFilial(bot, msg)
		user.State = StateAskFilial

	case StateAskFilial:
		user.Filial = ParseString(bot, msg, "филиал", assets.Filials[:])
		handle.HandleSelectCourse(bot, msg)
		user.State = StateAskCourse

	case StateAskCourse:

		user.Course = ParseString(bot, msg, "курс", assets.Courses[:])
		handle.HandleSelectFaculty(bot, msg)
		user.State = StateAskFaculty

	case StateAskFaculty:

		user.Faculty = ParseString(bot, msg, "факультет", assets.Fakultets[:])
		handle.HandleSelectCathedra(bot, msg)
		user.State = StateAskCathedra

	case StateAskCathedra:

		user.Cathedra = ParseString(bot, msg, "кафедра", assets.Cathedras[:])
		// handle.HandleSelectCathedra(bot, msg)
		user.State = StateConfirm
		filial := strings.Split(user.Filial, " ")[0]
		_, _ = bot.SendMessage(tu.Message(
			msg.Chat.ChatID(),
			fmt.Sprintf(
				"Выходит твой друг учится в %sом филиале на  %s, на  %s%s, верно? (Y/N)",
				filial[:len(filial)-2], user.Course, user.Faculty, user.Cathedra,
			),
		))
	case StateConfirm:

		if msg.Text == "Y" {
			_, _ = bot.SendMessage(tu.Message(
				msg.Chat.ChatID(),
				fmt.Sprintf("Thanks for your data!"),
			))
			user.State = StateSearch

		} else {
			handle.HandleSelectFilial(bot, msg)
			user.State = StateAskFilial
		}
	case StateSearch:
		SearchGroupUID(user.Filial, user.Course, user.Faculty, user.Cathedra)

	default:
		_, _ = bot.SendMessage(tu.Message(
			msg.Chat.ChatID(),
			fmt.Sprintf("Неизвестная команда"),
		))
		panic("unknown state")
	}
}
