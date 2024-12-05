package logic

import (
	"errors"

	"github.com/Robocotik/IT-s_Friends/assets/consts"
	"github.com/Robocotik/IT-s_Friends/internal/database"
	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/input"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"github.com/Robocotik/IT-s_Friends/internal/transport/handlers"

	"github.com/mymmrac/telego"
)

func FillObjectWithInfo(state *structures.State, bd database.IBd, bot *telego.Bot, msg telego.Message, identity *structures.Identity, isMe bool) {
	var err error
	switch *state {
	case structures.StateAskFilial:
		filials := bd.GetFilials(bot, msg)
		identity.Filial, err = input.ParseString(bot, msg, errors.New("филиал"), filials)
		if err != nil {
			handle.HandleSelectFilial(bd, bot, msg)
			break
		}
		handle.HandleSelectFaculty(bd, bot, msg, identity)
		*state = structures.StateAskFaculty

	case structures.StateAskFaculty:
		faculties := bd.GetFaculties(bot, msg, identity)
		identity.Faculty, err = input.ParseString(bot, msg, errors.New("факультет"), faculties)
		if err != nil {
			handle.HandleSelectFaculty(bd, bot, msg, identity)
			break
		}
		handle.HandleSelectCathedra(bd, bot, msg, identity)
		*state = structures.StateAskCathedra

	case structures.StateAskCathedra:
		cathedras := bd.GetCathedras(bot, msg, identity)
		identity.Cathedra, err = input.ParseString(bot, msg, errors.New("кафедра"), cathedras)
		if err != nil {
			handle.HandleSelectCathedra(bd, bot, msg, identity)
			break
		}
		*state = structures.StateAskCourse
		handle.HandleSelectCourse(bd, bot, msg, identity)

	case structures.StateAskCourse:
		courses := bd.GetCourses(bot, msg, identity)
		identity.Course, err = input.ParseString(bot, msg, errors.New("курс"), courses)
		if err != nil {
			handle.HandleSelectCourse(bd, bot, msg, identity)
			break
		}
		handle.HandleSelectGroup(bd, bot, msg, identity)
		*state = structures.StateAskGroup

	case structures.StateAskGroup:
		var groups = bd.GetGroups(bot, msg, identity)
		identity.Group, err = input.ParseString(bot, msg, errors.New("группа"), groups)
		if err != nil {
			handle.HandleSelectGroup(bd, bot, msg, identity)
			break
		}
		handle.HandleConfirm(bot, msg, identity, isMe)
		*state = structures.StateConfirm

	case structures.StateConfirm:
		_, err = input.ParseString(bot, msg, errors.New("ответ"), []string{consts.YES, consts.NO})
		if err != nil {
			handle.HandleConfirm(bot, msg, identity, isMe)
			break
		}
		if msg.Text == consts.YES {
			handle.HandleThankForData(bot, msg)
			identity.Uuid = bd.GetGroupByUID(bot, msg, identity)
			*state = structures.StateSearch

		} else {
			handle.HandleSelectFilial(bd, bot, msg)
			*state = structures.StateAskFilial
		}

	default:
		output.WriteMessage(bot, msg, "Неизвестная команда")
		panic("unknown state")
	}

}
