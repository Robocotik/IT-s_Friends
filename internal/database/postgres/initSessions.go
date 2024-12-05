package postgres

import (
	"fmt"

	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
)

func (psql Postgres) InitSessions(sessions *map[int64]*structures.User) {
	users, err := psql.GetAllIds()
	if err != nil {
		fmt.Println("Ошибка при получении всех пользователей ", err)
	}
	for _, userId := range users {
		(*sessions)[userId] = &structures.User{
			Id:     userId,
			Exists: true,
		}
	}

}
