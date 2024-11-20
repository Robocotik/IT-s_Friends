package server

import (
	// "Friends/src/components/structures"
	"context"
	// "encoding/json"
	"fmt"
	"os"
	// "strconv"

	"github.com/jackc/pgx/v5"
)

func SetUserId(conn *pgx.Conn, id int64) error {
	fmt.Sprintf("Я ДОБАВИЛ ID1: %s", id)
	// jsonn, _ := json.Marshal(structures.Fav{nickname, int64(idi)})
	// fmt.Println("У МЕНЯ JSON: ", jsonn)
	_, err := conn.Exec(context.Background(), "insert into users (id) values ($1) on conflict (id) do nothing", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}
	return nil
}
