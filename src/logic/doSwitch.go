package logic

import (
	"Friends/src/assets"
	keyboard "Friends/src/components/keyboards"
	"Friends/src/components/structures"
	"Friends/src/entities"
	"Friends/src/handlers"
	"Friends/src/utils"
	"Friends/src/utils/server"
	"errors"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var found_uid = ""
var ch_zn_selected = ""
var last_request = entities.Final_timetable{}

func DoSwitch(conn *pgx.Conn, user *structures.User, bot *telego.Bot, msg telego.Message) {
	var user_uid = strconv.Itoa(int(msg.Chat.ChatID().ID))
	var user_username = msg.Chat.ChatID().Username

	switch user.State {

	case structures.StateStart:
		server.SetUserId(conn, user_uid, user_username)
		handle.HandleStart(bot, msg)
		user.State = structures.StateDefault

	case structures.StateDefault:
		_ = utils.ParseString(bot, msg, errors.New("Погнали"), []string{"Погнали"})
		// if user.Filial == "-1" {
		// 	user.State = StateStart
		// 	break
		// }
		handle.HandleMenuStart(bot, msg)
		user.State = structures.StateStartMenu

	case structures.StateStartMenu:
		if msg.Text == structures.FIND_NEW_FRIENDS {
			handle.HandleSelectFilial(bot, msg)
			user.State = structures.StateAskFilial
		} else {
			favs, err := server.GetFavsFromId(conn, user_uid)
			utils.RiseError(bot, msg, err)
			utils.FuncWithKeyboard(bot, msg, func() []string {
				return utils.ShowFavs(favs)
			}, keyboard.CreateKeyboardReturnToSearch())
			user.State = structures.StateRedirectToStartSearch

		}

	case structures.StateAskFilial:
		filials := assets.GetFilials()
		user.Filial = utils.ParseString(bot, msg, errors.New("филиал"), filials)
		handle.HandleSelectFaculty(bot, msg, user.Filial)
		user.State = structures.StateAskFaculty

	case structures.StateAskFaculty:
		faculties := assets.GetFaculties(user.Filial)
		user.Faculty = utils.ParseString(bot, msg, errors.New("факультет"), faculties)
		handle.HandleSelectCathedra(bot, msg, user.Filial, user.Faculty)
		user.State = structures.StateAskCathedra

	case structures.StateAskCathedra:
		cathedras := assets.GetCathedras(user.Filial, user.Faculty)
		user.Cathedra = utils.ParseString(bot, msg, errors.New("кафедра"), cathedras)
		user.State = structures.StateAskCourse
		handle.HandleSelectCourse(bot, msg, user.Filial, user.Faculty, user.Cathedra)

	case structures.StateAskCourse:
		courses := assets.GetCourses(user.Filial, user.Faculty, user.Cathedra)
		user.Course = utils.ParseString(bot, msg, errors.New("курс"), courses)
		handle.HandleSelectGroup(bot, msg, user.Filial, user.Faculty, user.Course, user.Cathedra)
		user.State = structures.StateAskGroup

	case structures.StateAskGroup:
		var groups = assets.GetGroups(user.Filial, user.Course, user.Faculty, user.Cathedra)
		user.Group = utils.ParseString(bot, msg, errors.New("группа"), groups)
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

		found_uid = SearchGroupUID(bot, msg, user)
		last_request = DoRequest(bot, msg, found_uid)
		if len(last_request.Data.Schedule) != 0 { // проверка на наличие расписания
			handle.HandleGroupFound(bot, msg)
			user.State = structures.StateGroupFound
		} else {
			handle.HandleGroupNotFound(bot, msg)
			user.State = structures.StateRedirectToStartSearch
		}

	case structures.StateGroupFound: // Группа была найдена
		if msg.Text == structures.ADD_TO_FAVOURITE {
			handle.HandleSelectNickname(bot, msg)
			user.State = structures.StateAskNickname
		} else {
			user.State = structures.StateShowTimetable
			handle.HandleShowTimetable(bot, msg)
		}

	case structures.StateAskNickname:
		user.NickName = msg.Text
		handle.HandleAddToHavourite(bot, msg)
		id, err := strconv.Atoi(found_uid)
		utils.RiseError(bot, msg, err)
		user.Favourite = append(user.Favourite, structures.Fav{
			Nickname: user.NickName,
			Id:       int64(id),
		})
		server.AddIdToFavs(bot, msg, conn, user_uid, found_uid)
		user.State = structures.StateRedirectToStartSearch

	case structures.StateShowTimetable: // Вывод расписания
		ch_zn_selected := utils.ParseContainString(bot, msg, errors.New("Неизвестная четность недели"), []string{structures.Ch, structures.Zn})
		ShowTimetable(bot, msg, last_request, ch_zn_selected)
		user.State = structures.StateRedirectToStartSearch

	case structures.StateRedirectToStartSearch:
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
