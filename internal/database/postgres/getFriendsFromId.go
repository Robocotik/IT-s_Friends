package postgres

import (
	"context"

	"github.com/Robocotik/IT-s_Friends/internal/models/structures"

)

func (psql Postgres) GetFriendsFromId(userID int64) ([]structures.IFriendsShort, error) {
	var friends []structures.IFriendsShort

	rows, err := psql.Conn.Query(context.Background(),
		"SELECT f.nickname, f.group_title FROM user_friend uf JOIN friends f ON uf.friend_id = f.friend_id WHERE uf.user_id = $1", userID)
	if err != nil {
		return []structures.IFriendsShort{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var nickname string
		var group_title string
		if err := rows.Scan(&nickname, &group_title); err != nil {
			return nil, err
		}
		friends = append(friends, structures.IFriendsShort{Nickname: nickname, Group_title: group_title})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return friends, nil
}
