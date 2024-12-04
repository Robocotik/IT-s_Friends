package structures

import "github.com/Robocotik/IT-s_Friends/internal/models/entities"

type AskedFriend struct {
	Identity Identity
	NickName string
	State State
	Request  entities.Final_timetable
}

type IFriendsShort struct{
	Nickname string
	Group_title string
}
