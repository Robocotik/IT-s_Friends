package postgres

import (
	"context"
)

func (psql Postgres) GetAllIds() ([]int64, error) {
	var res []int64
	rows, err := psql.Conn.Query(context.Background(),
		"SELECT id FROM users")
	if err != nil {
		return []int64{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		res = append(res, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
