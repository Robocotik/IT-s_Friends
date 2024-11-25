package structures

import "Friends/src/entities"

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
