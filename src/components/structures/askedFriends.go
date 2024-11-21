package structures

import "Friends/src/entities"

type AskedFriend struct {
	Filial   string
	Faculty  string
	Course   string
	Cathedra string
	Group    string
	NickName string
	Uuid     string
	Request  entities.Final_timetable
}

type IFriendsShort struct{
	Nickname string
	Group_title string
}
