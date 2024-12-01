package logic

import (
	"errors"
	"github.com/Robocotik/IT-s_Friends/assets"
	"github.com/Robocotik/IT-s_Friends/assets/consts"
	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/input"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	"github.com/Robocotik/IT-s_Friends/internal/transport/handlers"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func FillObjectWithInfo(state *structures.State, conn *pgx.Conn, bot *telego.Bot, msg telego.Message, identity *structures.Identity, isMe bool) {
	var err error
	switch *state {
	case structures.StateAskFilial:
		filials := assets.GetFilials(conn, bot, msg)
		identity.Filial, err = input.ParseString(bot, msg, errors.New("филиал"), filials)
		if err != nil {
			handle.HandleSelectFilial(conn, bot, msg)
			break
		}
		handle.HandleSelectFaculty(conn, bot, msg, identity)
		*state = structures.StateAskFaculty

	case structures.StateAskFaculty:
		faculties := assets.GetFaculties(conn, bot, msg, identity)
		identity.Faculty, err = input.ParseString(bot, msg, errors.New("факультет"), faculties)
		if err != nil {
			handle.HandleSelectFaculty(conn, bot, msg, identity)
			break
		}
		handle.HandleSelectCathedra(conn, bot, msg, identity)
		*state = structures.StateAskCathedra

	case structures.StateAskCathedra:
		cathedras := assets.GetCathedras(conn, bot, msg, identity)
		identity.Cathedra, err = input.ParseString(bot, msg, errors.New("кафедра"), cathedras)
		if err != nil {
			handle.HandleSelectCathedra(conn, bot, msg, identity)
			break
		}
		*state = structures.StateAskCourse
		handle.HandleSelectCourse(conn, bot, msg, identity)

	case structures.StateAskCourse:
		courses := assets.GetCourses(conn, bot, msg, identity)
		identity.Course, err = input.ParseString(bot, msg, errors.New("курс"), courses)
		if err != nil {
			handle.HandleSelectCourse(conn, bot, msg, identity)
			break
		}
		handle.HandleSelectGroup(conn, bot, msg, identity)
		*state = structures.StateAskGroup

	case structures.StateAskGroup:
		var groups = assets.GetGroups(conn, bot, msg, identity)
		identity.Group, err = input.ParseString(bot, msg, errors.New("группа"), groups)
		if err != nil {
			handle.HandleSelectGroup(conn, bot, msg, identity)
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
			identity.Uuid = SearchGroupUID(bot, msg, conn, identity)
			*state = structures.StateSearch

		} else {
			handle.HandleSelectFilial(conn, bot, msg)
			*state = structures.StateAskFilial
		}

	default:
		output.WriteMessage(bot, msg, "Неизвестная команда")
		panic("unknown state")
	}

}
