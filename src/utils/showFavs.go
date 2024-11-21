package utils

import (
	"Friends/src/components/structures"
	"errors"
)

func ShowFavs(favs []structures.IFriendsShort) (string, error) {
	if len(favs) == 0 {
		return "", errors.New("no friends")
	}
	var res string
	for _, friend := range favs {
		res += "🐘 " + friend.Nickname + " | " + friend.Group_title + "\n"
	}
	return res, nil
}
