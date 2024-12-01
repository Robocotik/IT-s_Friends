package output

import (
	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"errors"
)

func ShowFavs(favs []structures.IFriendsShort) (string, error) {
	if len(favs) == 0 {
		return "У вас пока нет друзей... 😔", errors.New("no friends")
	}
	var res string
	for _, friend := range favs {
		res += "🐘 " + friend.Nickname + " | " + friend.Group_title + "\n"
	}
	return res, nil
}
