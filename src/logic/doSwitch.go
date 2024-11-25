package logic

import (
	"Friends/src/assets"
	keyboard "Friends/src/components/keyboards"
	"Friends/src/components/structures"
	"Friends/src/handlers"
	"Friends/src/utils"
	// "Friends/src/utils/bd"
	"Friends/src/utils/server"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

var ch_zn_selected = ""

func DoSwitch(conn *pgx.Conn, user *structures.User, friend *structures.AskedFriend, bot *telego.Bot, msg telego.Message) {
	var err error
	switch user.State {
	case structures.StateStart:
		// bd.ParseAllSchdule(conn, bot, msg)
		handle.HandleStart(bot, msg)
		user.State = structures.StateDefault

	case structures.StateDefault:
		_, err = utils.ParseString(bot, msg, errors.New("ответ"), []string{"Погнали"})
		if err != nil {
			handle.HandleStart(bot, msg)
			break
		}
		handle.HandleMenuStart(bot, msg)
		user.State = structures.StateStartMenu

	case structures.StateStartMenu:
		server.AddUserId(bot, msg, conn, msg.Chat.ChatID().ID, msg.From.Username)
		_, err = utils.ParseString(bot, msg, errors.New("ответ"), []string{structures.FIND_NEW_FRIENDS, structures.SHOW_FRIENDS})
		if err != nil {
			handle.HandleMenuStart(bot, msg)
			break
		}
		if msg.Text == structures.FIND_NEW_FRIENDS {
			handle.HandleSelectFilial(conn, bot, msg)
			user.State = structures.StateAskFilial
		} else {
			favs, err := server.GetFriendsFromId(conn, msg.Chat.ChatID().ID)
			utils.RiseError(bot, msg, err)
			utils.FuncWithKeyboard(bot, msg, func() (string, error) {
				return utils.ShowFavs(favs)
			}, keyboard.CreateKeyboardReturnToSearch())
			user.State = structures.StateRedirectToStartSearch

		}

	case structures.StateAskFilial:
		filials := assets.GetFilials(conn, bot, msg)
		friend.Filial, err = utils.ParseString(bot, msg, errors.New("филиал"), filials)
		if err != nil {
			handle.HandleSelectFilial(conn, bot, msg)
			break
		}
		handle.HandleSelectFaculty(conn, bot, msg, friend)
		user.State = structures.StateAskFaculty

	case structures.StateAskFaculty:
		faculties := assets.GetFaculties(conn, bot, msg, friend)
		friend.Faculty, err = utils.ParseString(bot, msg, errors.New("факультет"), faculties)
		if err != nil {
			handle.HandleSelectFaculty(conn, bot, msg, friend)
			break
		}
		handle.HandleSelectCathedra(conn, bot, msg, friend)
		user.State = structures.StateAskCathedra

	case structures.StateAskCathedra:
		cathedras := assets.GetCathedras(conn, bot, msg, friend)
		friend.Cathedra, err = utils.ParseString(bot, msg, errors.New("кафедра"), cathedras)
		if err != nil {
			handle.HandleSelectCathedra(conn, bot, msg, friend)
			break
		}
		user.State = structures.StateAskCourse
		handle.HandleSelectCourse(conn, bot, msg, friend)

	case structures.StateAskCourse:
		courses := assets.GetCourses(conn, bot, msg, friend)
		friend.Course, err = utils.ParseString(bot, msg, errors.New("курс"), courses)
		if err != nil {
			handle.HandleSelectCourse(conn, bot, msg, friend)
			break
		}
		handle.HandleSelectGroup(conn, bot, msg, friend)
		user.State = structures.StateAskGroup

	case structures.StateAskGroup:
		var groups = assets.GetGroups(conn, bot, msg, friend)
		friend.Group, err = utils.ParseString(bot, msg, errors.New("группа"), groups)
		if err != nil {
			handle.HandleSelectGroup(conn, bot, msg, friend)
			break
		}
		user.State = structures.StateConfirm
		handle.HandleConfirm(bot, msg, friend)

	case structures.StateConfirm:
		_, err = utils.ParseString(bot, msg, errors.New("ответ"), []string{structures.YES, structures.NO})
		if err != nil {
			handle.HandleConfirm(bot, msg, friend)
			break
		}
		if msg.Text == structures.YES {
			handle.HandleThankForData(bot, msg)
			user.State = structures.StateSearch

		} else {
			handle.HandleSelectFilial(conn, bot, msg)
			user.State = structures.StateAskFilial
		}

	case structures.StateSearch:

		friend.Uuid = SearchGroupUID(bot, msg, conn, friend)
		friend.Request = DoRequest(bot, msg, friend.Uuid)
		if len(friend.Request.Data.Schedule) != 0 { // проверка на наличие расписания
			handle.HandleGroupFound(bot, msg)
			user.State = structures.StateGroupFound
		} else {
			handle.HandleGroupNotFound(bot, msg)
			user.State = structures.StateRedirectToStartSearch
		}

	case structures.StateGroupFound: // Группа была найдена
		_, err = utils.ParseString(bot, msg, errors.New("ответ"), []string{structures.ADD_TO_FAVOURITE, structures.SHOW_SCHEDULE})
		if err != nil {
			handle.HandleConfirm(bot, msg, friend)
			break
		}
		if msg.Text == structures.ADD_TO_FAVOURITE {
			handle.HandleSelectNickname(bot, msg)
			user.State = structures.StateAskNickname
		} else {
			user.State = structures.StateShowTimetable
			handle.HandleShowTimetable(bot, msg)
		}

	case structures.StateAskNickname:
		friend.NickName = msg.Text
		friend_id, err := server.AddFriend(bot, msg, conn, friend)
		if err != nil {
			if err.Error() == "too long" {
				utils.WriteMessage(bot, msg, "К сожалению, имя друга превышает 40 символов, попробуй еще раз :)")
				handle.HandleSelectNickname(bot, msg)
				break
			}
		} else {
			server.AddConnection(bot, msg, conn, msg.Chat.ChatID().ID, friend_id)
			handle.HandleAddToHavourite(bot, msg)
		}
		user.State = structures.StateRedirectToStartSearch

	case structures.StateShowTimetable: // Вывод расписания
		ch_zn_selected := utils.ParseContainString(bot, msg, errors.New("Неизвестная четность недели"), []string{structures.Ch, structures.Zn})
		ShowTimetable(bot, msg, friend.Request, ch_zn_selected)
		user.State = structures.StateRedirectToStartSearch

	case structures.StateRedirectToStartSearch:
		handle.HandleMenuStart(bot, msg)
		user.State = structures.StateStartMenu

	default:
		utils.WriteMessage(bot, msg, "Неизвестная команда")
		panic("unknown state")
	}
}
