package postgres

import "github.com/jackc/pgx/v5"

type Reader interface {
	GetFriends()
	GetFriendsFromId()
}

type Updater interface {
	UpdateUser()
}

type Inserter interface {
	AddConnection()
	AddFriend()
	AddUserById()

}

type BD interface {
	Reader
	Updater
	Inserter
}

type postgres struct {
	conn *pgx.Conn
	Reader
	Updater
	Inserter
}
