package logic

import (
	"context"
	"errors"

	"github.com/Robocotik/IT-s_Friends/assets/consts"
	keyboard "github.com/Robocotik/IT-s_Friends/assets/keyboards"
	"github.com/Robocotik/IT-s_Friends/internal/database"
	errorsCustom "github.com/Robocotik/IT-s_Friends/internal/models/errors"
	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/input"
	"github.com/Robocotik/IT-s_Friends/internal/services/output"
	output_friends "github.com/Robocotik/IT-s_Friends/internal/services/output/friends"
	output_schedule "github.com/Robocotik/IT-s_Friends/internal/services/output/schedule"
	"github.com/Robocotik/IT-s_Friends/internal/services/utils"
	handle "github.com/Robocotik/IT-s_Friends/internal/transport/handlers"
	"github.com/mymmrac/telego"
)

var ch_zn_selected = ""

func DoSwitch(ctx context.Context, bd database.IBd, user *structures.User, friend *structures.AskedFriend, bot *telego.Bot, msg telego.Message) {
	var err error
	switch user.State {
	case structures.StateStart:
		handle.HandleStart(bot, msg)
		user.State = structures.StateDefault

	case structures.StateDefault:
		_, err = input.ParseString(bot, msg, errors.New("ответ"), []string{"Погнали"})
		if err != nil {
			handle.HandleStart(bot, msg)
			break
		}
		if user.Exists {
			handle.HandleMenuStart(bot, msg)
			user.State = structures.StateStartMenu
		} else {
			handle.HandleAddMe(bot, msg)
			user.State = structures.UserNotExists
		}
	case structures.UserNotExists:
		handle.HandleSelectFilial(bd, bot, msg)
		user.State = structures.StateAskForMe

	case structures.StateAskForMe:
		FillObjectWithInfo(&user.StateFilling, bd, bot, msg, &user.Identity, true)
		if user.StateFilling == structures.StateSearch {
			handle.HandleMenuStart(bot, msg)
			user.State = structures.StateStartMenu
			user.Exists = true
			bd.UpdateUser(bot, msg.Chat.ID, user.Identity, msg.Chat.ID)
		}
	case structures.StateStartMenu:
		_, err = input.ParseString(bot, msg, errors.New("ответ"), []string{consts.FIND_NEW_FRIENDS, consts.SHOW_FRIENDS, consts.SET_NOTIFICATIONS})
		if err != nil {
			handle.HandleMenuStart(bot, msg)
			break
		}
		switch msg.Text {
		case consts.FIND_NEW_FRIENDS:
			handle.HandleSelectFilial(bd, bot, msg)
			user.State = structures.StateAskForFriend

		case consts.SET_NOTIFICATIONS:
			handle.HandleSetNotifications(bot, msg)
			user.State = structures.StateSetNotifications

		default:
			favs, err := bd.GetFriendsFromId(msg.Chat.ChatID().ID)
			output.RiseError(bot, msg.Chat.ID, err)
			utils.FuncWithKeyboard(bot, msg, func() (string, error) {
				return output_friends.ShowFavs(favs)
			}, keyboard.CreateKeyboardReturnToSearch())
			user.State = structures.StateRedirectToStartSearch
		}

	case structures.StateSetNotifications:
		val, err2 := input.ParseString(bot, msg, errors.New("вариант"), []string{consts.CUSTOM_TIME, consts.H1_BEFORE, consts.H2_BEFORE, consts.H3_BEFORE})
		output.RiseError(bot, msg.Chat.ID, err2)
		if err2 != nil {
			handle.HandleSetNotifications(bot, msg)
			break
		}
		if val == consts.CUSTOM_TIME {
			user.State = structures.StateSetCustomNotification
			handle.HandleSetCustomNotification(bot, msg)

		} else {
			handle.HandleNotificationCreated(bot, msg)
			user.State = structures.StateRedirectToStartSearch
			// add server notification
		}

	case structures.StateSetCustomNotification:
		user.NotifyPeriod, err = input.CheckPeriod(msg.Text)
		if err != nil {
			output.WriteMessage(bot, msg.Chat.ID, "Неверный формат данных")
			handle.HandleSetCustomNotification(bot, msg)
		}
		// add server notification

	case structures.StateAskForFriend:
		FillObjectWithInfo(&user.Friend.State, bd, bot, msg, &friend.Identity, false)
		if user.Friend.State == structures.StateSearch {
			user.State = structures.StateSearch
			user.Friend.State = structures.StateAskFilial
		}

	case structures.StateSearch:
		friend.Identity.Uuid = bd.GetGroupByUID(bot, msg.Chat.ID, &friend.Identity)
		friend.Request = DoRequest(bot, msg, friend.Identity.Uuid)
		if len(friend.Request.Data.Schedule) != 0 { // проверка на наличие расписания
			handle.HandleGroupFound(bot, msg)
			user.State = structures.StateGroupFound
		} else {
			handle.HandleGroupNotFound(bot, msg)
			user.State = structures.StateRedirectToStartSearch
		}

	case structures.StateGroupFound: // Группа была найдена
		_, err = input.ParseString(bot, msg, errors.New("ответ"), []string{consts.ADD_TO_FAVOURITE, consts.SHOW_SCHEDULE})
		if err != nil {
			handle.HandleConfirm(bot, msg, &friend.Identity, false)
			break
		}
		if msg.Text == consts.ADD_TO_FAVOURITE {
			handle.HandleSelectNickname(bot, msg)
			user.State = structures.StateAskNickname
		} else {
			user.State = structures.StateShowTimetable
			handle.HandleShowTimetable(bot, msg)
		}

	case structures.StateAskNickname:
		friend.NickName = msg.Text
		friend_id, err := bd.AddFriend(bot, msg, friend)
		if err != nil {
			if err.Error() == errorsCustom.ErrTooLongMessage_23514 {
				output.WriteMessage(bot, msg.Chat.ID, errorsCustom.ErrTooLongMessage_23514)
				handle.HandleSelectNickname(bot, msg)
				break
			}
		} else {
			bd.AddConnection(ctx, bot, msg.Chat.ID, msg.Chat.ChatID().ID, friend_id)
			handle.HandleAddToHavourite(bot, msg)
		}
		user.State = structures.StateRedirectToStartSearch

	case structures.StateShowTimetable: // Вывод расписания
		ch_zn_selected, err := input.ParseContainString(bot, msg, []string{consts.Ch, consts.Zn})
		output.RiseError(bot, msg.Chat.ID, err)
		keyboard := keyboard.CreateKeyboardShowTimetable()
		output_schedule.ShowTimetable(bot, msg, keyboard, friend.Request, ch_zn_selected)
		user.State = structures.StateRedirectToStartSearch

	case structures.StateRedirectToStartSearch:
		handle.HandleMenuStart(bot, msg)
		user.State = structures.StateStartMenu

	default:
		output.WriteMessage(bot, msg.Chat.ID, "Неизвестная команда")
		panic("unknown state")
	}
}
