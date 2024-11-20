package utils

import "Friends/src/components/structures"

func ShowFavs(favs []structures.Fav) []string {
	if len(favs) == 0 {
		return []string{}
	}

	res := []string{}
	for _, fav := range favs {
		res = append(res, fav.Nickname)
	}

	return res
}
