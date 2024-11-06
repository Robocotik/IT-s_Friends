package utils

import (
	"Friends/src/assets"
	"Friends/src/components/structures"
	"Friends/src/handlers"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	StateStart structures.State = iota
	StateDefault
	StateAskCourse
	StateAskFaculty
	StateAskCathedra
	StateConfirm
	StateSearch
)

func DoSwitch(user *structures.User, bot *telego.Bot, msg telego.Message) {
	switch user.State {

	case StateStart:
		SearchGroupUID()
		msg.Text = ""
		handle.HandleStart(bot, msg)
		user.State = StateDefault
		msg.Text = ""

	case StateDefault:
		user.Course = ParseString(bot, msg, "Погнали", []string{"Погнали"})
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
		_, _ = bot.SendMessage(tu.Message(
			msg.Chat.ChatID(),
			fmt.Sprintf(
				"Выходит твой друг учится на  %s, на  %s%s, Верно? (Y/N)",
				user.Course, user.Faculty, user.Cathedra,
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
			_, _ = bot.SendMessage(tu.Message(
				msg.Chat.ChatID(),
				fmt.Sprintf("Ok, let's start again"),
			))

			user.State = StateStart
		}
	case StateSearch:
		SearchGroupUID()

	default:
		_, _ = bot.SendMessage(tu.Message(
			msg.Chat.ChatID(),
			fmt.Sprintf("Неизвестная команда"),
		))
		panic("unknown state")
	}
}
