package database

import (
	"context"

	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/mymmrac/telego"
)

type AssetsReader interface {
	GetCathedras(bot *telego.Bot, chatID int64, identity *structures.Identity) []string
	GetCourses(bot *telego.Bot, chatID int64, identity *structures.Identity) []string
	GetFaculties(bot *telego.Bot, chatID int64, identity *structures.Identity) []string
	GetFilials(bot *telego.Bot, chatID int64) []string
	GetGroups(bot *telego.Bot, chatID int64, identity *structures.Identity) []string
}

type Reader interface {
	GetAllIds() ([]int64, error)
	GetFriendsFromId(userID int64) ([]structures.IFriendsShort, error)
	GetGroupByUID(bot *telego.Bot, chatID int64, identity *structures.Identity) string
}

type Updater interface {
	UpdateUser(bot *telego.Bot, chatID int64, identity structures.Identity, id int64) error
}

type Inserter interface {
	AddConnection(ctx context.Context, bot *telego.Bot, chatID int64, user_id int64, friend_id int64) error
	AddFriend(bot *telego.Bot, msg telego.Message, friend *structures.AskedFriend) (int64, error)
	AddUserId(bot *telego.Bot, chatID int64, id int64, nickname string) error
}

type IBd interface {
	Reader
	Updater
	Inserter
	AssetsReader
	InitSessions(sessions *map[int64]*structures.User)
}

// func (bd *Database) NewDatabase() error {
// 	fmt.Println("Empty database implementation")
// 	return nil
// }

// func (bd *Database) GetFriendsFromId(userID int64) ([]structures.IFriendsShort, error) {
// 	fmt.Println("Empty database implementation")
// 	return []structures.IFriendsShort{}, nil
// }

// func (bd *Database) GetAllIds() ([]int64, error) {
// 	fmt.Println("Empty database implementation")
// 	return []int64{}, nil
// }

// func (bd *Database) UpdateUser(bot *telego.Bot, msg telego.Message, identity structures.Identity, id int64) error {
// 	fmt.Println("Empty database implementation")
// 	return nil
// }

// func (bd *Database) AddConnection(ctx context.Context, bot *telego.Bot, msg telego.Message, user_id int64, friend_id int64) error {
// 	fmt.Println("Empty database implementation")
// 	return nil
// }

// func (bd *Database) AddFriend(bot *telego.Bot, msg telego.Message, friend *structures.AskedFriend) (int64, error) {
// 	fmt.Println("Empty database implementation")
// 	return 0, nil
// }
// func (bd *Database) AddUserId(bot *telego.Bot, msg telego.Message, id int64, nickname string) error {
// 	fmt.Println("Empty database implementation")
// 	return nil
// }

// func (bd *Database) InitSessions(sessions *map[int64]*structures.User) {
// 	fmt.Println("Empty database implementation")
// }
