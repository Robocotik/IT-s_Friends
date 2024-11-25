package logic

import (
	"github.com/Robocotik/IT-s_Friends/src/assets"
	"github.com/Robocotik/IT-s_Friends/src/components/structures"
	handle "github.com/Robocotik/IT-s_Friends/src/handlers"
	"github.com/Robocotik/IT-s_Friends/src/utils"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func FillObjectWithInfo(state *structures.State, conn *pgx.Conn, bot *telego.Bot, msg telego.Message, identity *structures.Identity) {
	var err error
	fmt.Println("ЩАС НА state: ", *state)
	switch *state {
	case structures.StateAskFilial:
		filials := assets.GetFilials(conn, bot, msg)
		identity.Filial, err = utils.ParseString(bot, msg, errors.New("филиал"), filials)
		if err != nil {
			handle.HandleSelectFilial(conn, bot, msg)
			break
		}
		handle.HandleSelectFaculty(conn, bot, msg, identity)
		*state = structures.StateAskFaculty

	case structures.StateAskFaculty:
		faculties := assets.GetFaculties(conn, bot, msg, identity)
		identity.Faculty, err = utils.ParseString(bot, msg, errors.New("факультет"), faculties)
		if err != nil {
			handle.HandleSelectFaculty(conn, bot, msg, identity)
			break
		}
		handle.HandleSelectCathedra(conn, bot, msg, identity)
		*state = structures.StateAskCathedra

	case structures.StateAskCathedra:
		cathedras := assets.GetCathedras(conn, bot, msg, identity)
		identity.Cathedra, err = utils.ParseString(bot, msg, errors.New("кафедра"), cathedras)
		if err != nil {
			handle.HandleSelectCathedra(conn, bot, msg, identity)
			break
		}
		*state = structures.StateAskCourse
		handle.HandleSelectCourse(conn, bot, msg, identity)

	case structures.StateAskCourse:
		courses := assets.GetCourses(conn, bot, msg, identity)
		identity.Course, err = utils.ParseString(bot, msg, errors.New("курс"), courses)
		if err != nil {
			handle.HandleSelectCourse(conn, bot, msg, identity)
			break
		}
		handle.HandleSelectGroup(conn, bot, msg, identity)
		*state = structures.StateAskGroup

	case structures.StateAskGroup:
		var groups = assets.GetGroups(conn, bot, msg, identity)
		identity.Group, err = utils.ParseString(bot, msg, errors.New("группа"), groups)
		if err != nil {
			handle.HandleSelectGroup(conn, bot, msg, identity)
			break
		}
		handle.HandleConfirm(bot, msg, identity)
		*state = structures.StateConfirm

	case structures.StateConfirm:
		_, err = utils.ParseString(bot, msg, errors.New("ответ4"), []string{structures.YES, structures.NO})
		if err != nil {
			handle.HandleConfirm(bot, msg, identity)
			break
		}
		if msg.Text == structures.YES {
			handle.HandleThankForData(bot, msg)
			*state = structures.StateSearch

		} else {
			handle.HandleSelectFilial(conn, bot, msg)
			*state = structures.StateAskFilial
		}

	default:
		utils.WriteMessage(bot, msg, "Неизвестная команда")
		panic("unknown state")
	}

}
