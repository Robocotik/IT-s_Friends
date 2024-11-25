package logic

import (
	keyboard "github.com/Robocotik/IT-s_Friends/src/components/keyboards"
	"github.com/Robocotik/IT-s_Friends/src/components/structures"
	errorsCustom "github.com/Robocotik/IT-s_Friends/src/components/structures/errors"
	"github.com/Robocotik/IT-s_Friends/src/handlers"
	"github.com/Robocotik/IT-s_Friends/src/utils"

	// "Friends/src/utils/bd"
	"github.com/Robocotik/IT-s_Friends/src/utils/server"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	// "strconv"
)

var ch_zn_selected = ""

func DoSwitch(conn *pgx.Conn, user *structures.User, friend *structures.AskedFriend, bot *telego.Bot, msg telego.Message, exists bool) {
	var err error
	// utils.WriteMessage(bot, msg, strconv.FormatBool(exists))
	switch user.State {
	case structures.StateStart:
		// bd.ParseAllSchdule(conn, bot, msg)
		handle.HandleStart(bot, msg)
		user.State = structures.StateDefault

	case structures.StateDefault:
		_, err = utils.ParseString(bot, msg, errors.New("ответ1"), []string{"Погнали"})
		if err != nil {
			handle.HandleStart(bot, msg)
			break
		}
		if exists {
			handle.HandleMenuStart(bot, msg)
			user.State = structures.StateStartMenu
		} else {
			handle.HandleAddMe(bot, msg)
			user.State = structures.StateAskForMe
		}

	case structures.StateAskForMe:
		FillObjectWithInfo(&user.StateFilling, conn, bot, msg, &user.Identity)
		if user.StateFilling == structures.StateSearch {
			handle.HandleAskForMe(bot, msg)
			user.State = structures.StateStartMenu
			server.SetInfoForId(bot, msg, conn, user.Identity, msg.Chat.ID )
		}
	case structures.StateStartMenu:
		_, err = utils.ParseString(bot, msg, errors.New("ответ2"), []string{structures.FIND_NEW_FRIENDS, structures.SHOW_FRIENDS})
		if err != nil {
			handle.HandleMenuStart(bot, msg)
			break
		}
		if msg.Text == structures.FIND_NEW_FRIENDS {
			handle.HandleSelectFilial(conn, bot, msg)
			user.State = structures.StateAskForFriend
		} else {
			favs, err := server.GetFriendsFromId(conn, msg.Chat.ChatID().ID)
			utils.RiseError(bot, msg, err)
			utils.FuncWithKeyboard(bot, msg, func() (string, error) {
				return utils.ShowFavs(favs)
			}, keyboard.CreateKeyboardReturnToSearch())
			user.State = structures.StateRedirectToStartSearch

		}

	case structures.StateAskForFriend:
		FillObjectWithInfo(&user.Friend.State, conn, bot, msg, &friend.Identity)
		if user.Friend.State  == structures.StateSearch {
			user.State = structures.StateSearch
			user.Friend.State = structures.StateAskFilial
		}

	case structures.StateSearch:
		friend.Identity.Uuid = SearchGroupUID(bot, msg, conn, friend)
		friend.Request = DoRequest(bot, msg, friend.Identity.Uuid)
		if len(friend.Request.Data.Schedule) != 0 { // проверка на наличие расписания
			handle.HandleGroupFound(bot, msg)
			user.State = structures.StateGroupFound
		} else {
			handle.HandleGroupNotFound(bot, msg)
			user.State = structures.StateRedirectToStartSearch
		}

	case structures.StateGroupFound: // Группа была найдена
		_, err = utils.ParseString(bot, msg, errors.New("ответ3"), []string{structures.ADD_TO_FAVOURITE, structures.SHOW_SCHEDULE})
		if err != nil {
			handle.HandleConfirm(bot, msg, &friend.Identity)
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
			if err.Error() == errorsCustom.ErrTooLongMessage_23514 {
				utils.WriteMessage(bot, msg, errorsCustom.ErrTooLongMessage_23514)
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
